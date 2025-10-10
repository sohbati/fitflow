package auth

import (
	"iam-service/internal/user"
	"iam-service/pkg/jwt"
)

// Handler manages authentication-related HTTP handlers
type Handler struct {
	userService         user.Service
	userAuthService     UserAuthService
	authProviderService AuthProviderService
	jwtManager          *jwt.JWTManager
	GoogleAuthHandler   *GoogleAuthHandler
}

// NewHandler creates a new authentication handler
func NewHandler(userService user.Service, userAuthService UserAuthService, authProviderService AuthProviderService, jwtManager *jwt.JWTManager) *Handler {
	return &Handler{
		userService:         userService,
		userAuthService:     userAuthService,
		authProviderService: authProviderService,
		jwtManager:          jwtManager,
	}
}

// NewHandlerWithGoogle creates a new authentication handler with Google OAuth
func NewHandlerWithGoogle(userService user.Service, userAuthService UserAuthService, authProviderService AuthProviderService, jwtManager *jwt.JWTManager, googleConfig *GoogleOAuthConfig) *Handler {
	googleAuthHandler := NewGoogleAuthHandler(userService, userAuthService, authProviderService, jwtManager, googleConfig)
	return &Handler{
		userService:         userService,
		userAuthService:     userAuthService,
		authProviderService: authProviderService,
		jwtManager:          jwtManager,
		GoogleAuthHandler:   googleAuthHandler,
	}
}
