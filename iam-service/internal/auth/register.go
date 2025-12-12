package auth

import (
	"errors"
	"iam-service/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRequest represents user registration data
type RegisterRequest struct {
	Username    string `json:"username" binding:"required" example:"john_doe"`
	Email       string `json:"email" binding:"required,email" example:"john.doe@example.com"`
	Mobile      string `json:"mobile" binding:"required" example:"+1234567890"`
	DisplayName string `json:"display_name" binding:"required" example:"John Doe"`
	Password    string `json:"password" binding:"required" example:"securepassword123"`
	Country     string `json:"country" binding:"required" example:"US"`
	Role        string `json:"role,omitempty" example:"user"`
}

// Register godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "User registration data"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /register [post]
func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.HandleError(c, err, http.StatusBadRequest)
		return
	}

	// Use LocalAuthHandler for registration
	localAuthReq := LocalRegisterRequest{
		Username:    req.Username,
		Email:       req.Email,
		Mobile:      req.Mobile,
		DisplayName: req.DisplayName,
		Password:    req.Password,
		Country:     req.Country,
		Role:        req.Role,
	}

	// Create local auth handler
	localAuthHandler := NewLocalAuthHandler(h.userService, h.userAuthService, h.authProviderService, h.jwtManager)

	response, err := localAuthHandler.Register(localAuthReq)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "email already exists" ||
			err.Error() == "mobile number already exists" ||
			err.Error() == "username already exists" {
			status = http.StatusConflict
		}
		middleware.HandleError(c, err, status)
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse{
		Message: "User created successfully",
		Data:    response,
	})
}
