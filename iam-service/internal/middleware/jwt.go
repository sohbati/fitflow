package middleware

import (
	"errors"
	"net/http"
	"strings"

	"iam-service/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func JWTAuth(jwtManager *jwt.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			HandleError(c, errors.New("missing_token: Authorization header is required"), http.StatusUnauthorized)
			return
		}

		// Check if the header starts with "Bearer "
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			HandleError(c, errors.New("invalid_token_format: Authorization header must start with 'Bearer '"), http.StatusUnauthorized)
			return
		}

		// Validate the token
		claims, err := jwtManager.ValidateToken(tokenString)
		if err != nil {
			HandleError(c, errors.New("invalid_token: Invalid or expired token"), http.StatusUnauthorized)
			return
		}

		// Set user information in context
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)

		c.Next()
	}
}
