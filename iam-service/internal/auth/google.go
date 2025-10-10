package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"iam-service/internal/user"
	"iam-service/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// GoogleOAuthConfig holds Google OAuth configuration
type GoogleOAuthConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

// GoogleUserInfo represents Google user information
type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

// GoogleAuthHandler handles Google OAuth authentication
type GoogleAuthHandler struct {
	userService         user.Service
	userAuthService     UserAuthService
	authProviderService AuthProviderService
	jwtManager          *jwt.JWTManager
	config              *oauth2.Config
}

// NewGoogleAuthHandler creates a new Google OAuth handler
func NewGoogleAuthHandler(userService user.Service, userAuthService UserAuthService, authProviderService AuthProviderService, jwtManager *jwt.JWTManager, googleConfig *GoogleOAuthConfig) *GoogleAuthHandler {
	config := &oauth2.Config{
		ClientID:     googleConfig.ClientID,
		ClientSecret: googleConfig.ClientSecret,
		RedirectURL:  googleConfig.RedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return &GoogleAuthHandler{
		userService:         userService,
		userAuthService:     userAuthService,
		authProviderService: authProviderService,
		jwtManager:          jwtManager,
		config:              config,
	}
}

// GoogleLoginRequest represents the request for Google login
type GoogleLoginRequest struct {
	Code string `json:"code" binding:"required"`
}

// GoogleLoginResponse represents the response for Google login
type GoogleLoginResponse struct {
	Token string    `json:"token"`
	User  user.User `json:"user"`
}

// GoogleLogin handles Google OAuth login
// @Summary Google OAuth Login
// @Description Authenticate user with Google OAuth
// @Tags authentication
// @Accept json
// @Produce json
// @Param request body GoogleLoginRequest true "Google OAuth code"
// @Success 200 {object} GoogleLoginResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/google [post]
func (h *GoogleAuthHandler) GoogleLogin(c *gin.Context) {
	var req GoogleLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("Google OAuth: Invalid request format: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	fmt.Printf("Google OAuth: Received code: %s\n", req.Code)

	// Exchange code for token
	token, err := h.config.Exchange(context.Background(), req.Code)
	if err != nil {
		fmt.Printf("Google OAuth: Failed to exchange code for token: %v\n", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to exchange code for token"})
		return
	}

	fmt.Printf("Google OAuth: Successfully exchanged code for token\n")

	// Get user info from Google
	userInfo, err := h.getGoogleUserInfo(token.AccessToken)
	if err != nil {
		fmt.Printf("Google OAuth: Failed to get user info from Google: %v\n", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to get user info from Google"})
		return
	}

	fmt.Printf("Google OAuth: Got user info: %s (%s)\n", userInfo.Name, userInfo.Email)

	// Get Google auth provider
	googleProvider, err := h.authProviderService.GetProviderByName(AuthProviderGoogle)
	if err != nil {
		fmt.Printf("Google OAuth: Google auth provider not found: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Google auth provider not found"})
		return
	}

	fmt.Printf("Google OAuth: Found Google provider: %s\n", googleProvider.ID)

	// Check if user auth exists for this Google ID
	existingUserAuth, err := h.userAuthService.GetUserAuthByProviderAndProviderUserID(googleProvider.ID, userInfo.ID)
	if err != nil {
		fmt.Printf("Google OAuth: User auth doesn't exist, checking if user exists by email\n")
		// User auth doesn't exist, check if user exists by email
		existingUser, err := h.userService.GetUserByEmail(userInfo.Email)
		if err != nil {
			fmt.Printf("Google OAuth: User doesn't exist, creating new user\n")
			// User doesn't exist, create new user
			newUser := &user.User{
				Email:       userInfo.Email,
				DisplayName: userInfo.Name,
				AvatarURL:   userInfo.Picture,
			}

			createdUser, err := h.userService.CreateUserFromStruct(newUser)
			if err != nil {
				fmt.Printf("Google OAuth: Failed to create user: %v\n", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
				return
			}

			fmt.Printf("Google OAuth: Created user: %s\n", createdUser.ID)

			// Create user auth for Google
			userAuth := &UserAuth{
				UserID:         createdUser.ID,
				AuthProviderID: googleProvider.ID,
				ProviderUserID: userInfo.ID,
				ProviderEmail:  userInfo.Email,
				IsVerified:     true,
			}

			// Add provider data as JSON
			providerData := map[string]interface{}{
				"name":        userInfo.Name,
				"given_name":  userInfo.GivenName,
				"family_name": userInfo.FamilyName,
				"picture":     userInfo.Picture,
				"locale":      userInfo.Locale,
			}
			if data, err := json.Marshal(providerData); err == nil {
				userAuth.ProviderData = string(data)
			}

			_, err = h.userAuthService.CreateUserAuth(userAuth)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user auth"})
				return
			}

			// Generate JWT token
			jwtToken, err := h.jwtManager.GenerateToken(createdUser.ID.String(), createdUser.Email)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
				return
			}

			c.JSON(http.StatusOK, GoogleLoginResponse{
				Token: jwtToken,
				User:  *createdUser,
			})
			return
		}

		// User exists, create user auth for Google
		userAuth := &UserAuth{
			UserID:         existingUser.ID,
			AuthProviderID: googleProvider.ID,
			ProviderUserID: userInfo.ID,
			ProviderEmail:  userInfo.Email,
			IsVerified:     true,
		}

		// Add provider data as JSON
		providerData := map[string]interface{}{
			"name":        userInfo.Name,
			"given_name":  userInfo.GivenName,
			"family_name": userInfo.FamilyName,
			"picture":     userInfo.Picture,
			"locale":      userInfo.Locale,
		}
		if data, err := json.Marshal(providerData); err == nil {
			userAuth.ProviderData = string(data)
		}

		_, err = h.userAuthService.CreateUserAuth(userAuth)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user auth"})
			return
		}

		// Update user info if needed
		if existingUser.AvatarURL != userInfo.Picture || existingUser.DisplayName != userInfo.Name {
			existingUser.AvatarURL = userInfo.Picture
			existingUser.DisplayName = userInfo.Name
			_, err = h.userService.UpdateUser(existingUser)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
				return
			}
		}

		// Generate JWT token
		jwtToken, err := h.jwtManager.GenerateToken(existingUser.ID.String(), existingUser.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, GoogleLoginResponse{
			Token: jwtToken,
			User:  *existingUser,
		})
		return
	}

	// User auth exists, get the user
	existingUser, err := h.userService.GetUserByID(existingUserAuth.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	// Update user info if needed
	if existingUser.AvatarURL != userInfo.Picture || existingUser.DisplayName != userInfo.Name {
		existingUser.AvatarURL = userInfo.Picture
		existingUser.DisplayName = userInfo.Name
		_, err = h.userService.UpdateUser(existingUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}
	}

	// Update user auth with latest provider data
	providerData := map[string]interface{}{
		"name":        userInfo.Name,
		"given_name":  userInfo.GivenName,
		"family_name": userInfo.FamilyName,
		"picture":     userInfo.Picture,
		"locale":      userInfo.Locale,
	}
	if data, err := json.Marshal(providerData); err == nil {
		existingUserAuth.ProviderData = string(data)
	}
	existingUserAuth.IsVerified = true
	h.userAuthService.UpdateUserAuth(existingUserAuth)

	// Generate JWT token
	jwtToken, err := h.jwtManager.GenerateToken(existingUser.ID.String(), existingUser.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, GoogleLoginResponse{
		Token: jwtToken,
		User:  *existingUser,
	})
}

// getGoogleUserInfo fetches user information from Google API
func (h *GoogleAuthHandler) getGoogleUserInfo(accessToken string) (*GoogleUserInfo, error) {
	url := fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", accessToken)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user info, status: %d", resp.StatusCode)
	}

	var userInfo GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}

// generateUsernameFromEmail generates a username from email
func (h *GoogleAuthHandler) generateUsernameFromEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) > 0 {
		return parts[0]
	}
	return email
}

// GetGoogleAuthURL returns the Google OAuth authorization URL
// @Summary Get Google OAuth URL
// @Description Get Google OAuth authorization URL
// @Tags authentication
// @Produce json
// @Success 200 {object} map[string]string
// @Router /auth/google/url [get]
func (h *GoogleAuthHandler) GetGoogleAuthURL(c *gin.Context) {
	url := h.config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.JSON(http.StatusOK, gin.H{"url": url})
}
