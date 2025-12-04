package auth

import (
	"iam-service/internal/session"
	"iam-service/internal/user"
	"iam-service/pkg/jwt"
)

// Handler manages authentication-related HTTP handlers
type Handler struct {
	userService         user.Service
	userAuthService     UserAuthService
	authProviderService AuthProviderService
	jwtManager          *jwt.JWTManager
	sessionService      session.Service
	GoogleAuthHandler   *GoogleAuthHandler
}

// NewHandler creates a new authentication handler
func NewHandler(userService user.Service, userAuthService UserAuthService, authProviderService AuthProviderService, jwtManager *jwt.JWTManager, sessionService session.Service) *Handler {
	return &Handler{
		userService:         userService,
		userAuthService:     userAuthService,
		authProviderService: authProviderService,
		jwtManager:          jwtManager,
		sessionService:      sessionService,
	}
}

// NewHandlerWithGoogle creates a new authentication handler with Google OAuth
func NewHandlerWithGoogle(userService user.Service, userAuthService UserAuthService, authProviderService AuthProviderService, jwtManager *jwt.JWTManager, sessionService session.Service, googleConfig *GoogleOAuthConfig) *Handler {
	googleAuthHandler := NewGoogleAuthHandler(userService, userAuthService, authProviderService, jwtManager, sessionService, googleConfig)
	return &Handler{
		userService:         userService,
		userAuthService:     userAuthService,
		authProviderService: authProviderService,
		jwtManager:          jwtManager,
		sessionService:      sessionService,
		GoogleAuthHandler:   googleAuthHandler,
	}
}
