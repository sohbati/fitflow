package auth

import (
	"iam-service/internal/middleware"
	"iam-service/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginRequest represents user login credentials
type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"john_doe"`
	Password string `json:"password" binding:"required" example:"securepassword123"`
}

// AuthResponse represents successful authentication response
type AuthResponse struct {
	Token string     `json:"token"`
	User  *user.User `json:"user"`
}

// Login godoc
// @Summary Authenticate user
// @Description Login user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login credentials"
// @Success 200 {object} AuthResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /login [post]
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.HandleError(c, err, http.StatusBadRequest)
		return
	}

	// Use LocalAuthHandler for login
	localAuthReq := LocalAuthRequest{
		Username: req.Username,
		Password: req.Password,
	}

	// Create local auth handler
	localAuthHandler := NewLocalAuthHandler(h.userService, h.userAuthService, h.authProviderService, h.jwtManager)

	response, err := localAuthHandler.Login(localAuthReq)
	if err != nil {
		middleware.HandleError(c, errors.New("authentication_failed"), http.StatusUnauthorized)
		return
	}

	c.JSON(http.StatusOK, AuthResponse{
		Token: response.Token,
		User:  &response.User,
	})
}
