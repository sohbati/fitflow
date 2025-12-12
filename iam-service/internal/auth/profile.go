package auth

import (
	"errors"
	"iam-service/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetMe godoc
// @Summary Get current user info
// @Description Get information about the currently authenticated user
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} SuccessResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /me [get]
func (h *Handler) GetMe(c *gin.Context) {
	userIDStr, exists := c.Get("user_id")
	if !exists {
		middleware.HandleError(c, errors.New("unauthorized: User ID not found in context"), http.StatusUnauthorized)
		return
	}

	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		middleware.HandleError(c, err, http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		middleware.HandleError(c, err, http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Message: "User information retrieved successfully",
		Data:    user,
	})
}
