package auth

import (
	"encoding/json"
	"errors"
	"iam-service/internal/user"
	"iam-service/pkg/crypto"
	"iam-service/pkg/jwt"
)

// LocalAuthRequest represents the request for local authentication
type LocalAuthRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LocalRegisterRequest represents the request for local registration
type LocalRegisterRequest struct {
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Mobile      string `json:"mobile" binding:"required"`
	DisplayName string `json:"display_name" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Country     string `json:"country" binding:"required"`
	Role        string `json:"role,omitempty"`
}

// LocalAuthResponse represents the response for local authentication
type LocalAuthResponse struct {
	Token string    `json:"token"`
	User  user.User `json:"user"`
}

// LocalAuthHandler handles local username/password authentication
type LocalAuthHandler struct {
	userService         user.Service
	userAuthService     UserAuthService
	authProviderService AuthProviderService
	jwtManager          *jwt.JWTManager
}

// NewLocalAuthHandler creates a new local authentication handler
func NewLocalAuthHandler(userService user.Service, userAuthService UserAuthService, authProviderService AuthProviderService, jwtManager *jwt.JWTManager) *LocalAuthHandler {
	return &LocalAuthHandler{
		userService:         userService,
		userAuthService:     userAuthService,
		authProviderService: authProviderService,
		jwtManager:          jwtManager,
	}
}

// Register handles local user registration
func (h *LocalAuthHandler) Register(req LocalRegisterRequest) (*LocalAuthResponse, error) {
	// Check if email already exists
	existingUser, _ := h.userService.GetUserByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := crypto.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Create user
	newUser := &user.User{
		Email:       req.Email,
		DisplayName: req.DisplayName,
		Country:     req.Country,
		IsActive:    true,
	}

	createdUser, err := h.userService.CreateUserFromStruct(newUser)
	if err != nil {
		return nil, err
	}

	// Get local auth provider
	localProvider, err := h.authProviderService.GetProviderByName(AuthProviderLocal)
	if err != nil {
		return nil, errors.New("local auth provider not found")
	}

	// Create user auth for local authentication
	userAuth := &UserAuth{
		UserID:         createdUser.ID,
		AuthProviderID: localProvider.ID,
		ProviderEmail:  req.Email,
		Username:       req.Username,
		IsVerified:     true,
	}

	// Store password and mobile in provider data
	providerData := map[string]interface{}{
		"password": hashedPassword,
		"mobile":   req.Mobile,
	}
	if data, err := json.Marshal(providerData); err == nil {
		userAuth.ProviderData = string(data)
	}

	_, err = h.userAuthService.CreateUserAuth(userAuth)
	if err != nil {
		return nil, err
	}

	// Generate JWT token
	jwtToken, err := h.jwtManager.GenerateToken(createdUser.ID.String(), createdUser.Email)
	if err != nil {
		return nil, err
	}

	return &LocalAuthResponse{
		Token: jwtToken,
		User:  *createdUser,
	}, nil
}

// Login handles local user authentication
func (h *LocalAuthHandler) Login(req LocalAuthRequest) (*LocalAuthResponse, error) {
	// Get local auth provider
	localProvider, err := h.authProviderService.GetProviderByName(AuthProviderLocal)
	if err != nil {
		return nil, errors.New("local auth provider not found")
	}

	// Find user auth by username
	matchingUserAuth, err := h.userAuthService.GetUserAuthByProviderAndUsername(localProvider.ID, req.Username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Get the user
	existingUser, err := h.userService.GetUserByID(matchingUserAuth.UserID)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Verify password
	var providerData map[string]interface{}
	if err := json.Unmarshal([]byte(matchingUserAuth.ProviderData), &providerData); err != nil {
		return nil, errors.New("invalid credentials")
	}

	hashedPassword, ok := providerData["password"].(string)
	if !ok {
		return nil, errors.New("invalid credentials")
	}

	if !crypto.VerifyPassword(req.Password, hashedPassword) {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT token
	jwtToken, err := h.jwtManager.GenerateToken(existingUser.ID.String(), existingUser.Email)
	if err != nil {
		return nil, err
	}

	return &LocalAuthResponse{
		Token: jwtToken,
		User:  *existingUser,
	}, nil
}
