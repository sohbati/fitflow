package handler

import (
	"fitflow-business/internal/model"
	"fitflow-business/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ProfileHandler handles HTTP requests for profile-related operations
type ProfileHandler struct {
	profileService service.ProfileService
}

// NewProfileHandler creates a new profile handler
func NewProfileHandler(profileService service.ProfileService) *ProfileHandler {
	return &ProfileHandler{
		profileService: profileService,
	}
}

// CreateProfileRequest represents the request to create a profile
type CreateProfileRequest struct {
	UserID     string `json:"user_id" binding:"required"`
	Type       string `json:"type" binding:"required"`
	PersonID   int64  `json:"person_id" binding:"required"`
	GymOwnerID *int64 `json:"gym_owner_id,omitempty"`
	TrainerID  *int64 `json:"trainer_id,omitempty"`
	TraineeID  *int64 `json:"trainee_id,omitempty"`
	IsActive   *bool  `json:"is_active,omitempty"`
}

// CreateProfile handles POST /profiles
func (h *ProfileHandler) CreateProfile(c *gin.Context) {
	var req CreateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	// Parse user ID
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	// Validate profile type
	profileType := model.ProfileType(req.Type)
	if !model.IsValidProfileType(profileType) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile type. Must be 'gym_owner', 'trainer', or 'trainee'"})
		return
	}

	// Create profile
	profile := &model.Profile{
		UserID:     userID,
		Type:       profileType,
		PersonID:   req.PersonID,
		GymOwnerID: req.GymOwnerID,
		TrainerID:  req.TrainerID,
		TraineeID:  req.TraineeID,
	}

	if req.IsActive != nil {
		profile.IsActive = *req.IsActive
	} else {
		profile.IsActive = true
	}

	if err := h.profileService.CreateProfile(c.Request.Context(), profile); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Profile created successfully",
		"profile": profile,
	})
}

// GetProfile handles GET /profiles/:id
func (h *ProfileHandler) GetProfile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile ID"})
		return
	}

	profile, err := h.profileService.GetProfileByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	c.JSON(http.StatusOK, profile)
}

// GetProfilesByUserID handles GET /profiles/user/:user_id
func (h *ProfileHandler) GetProfilesByUserID(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	profiles, err := h.profileService.GetProfilesByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"profiles": profiles,
		"count":    len(profiles),
	})
}

// GetProfileByType handles GET /profiles/user/:user_id/type/:type
func (h *ProfileHandler) GetProfileByType(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	profileTypeStr := c.Param("type")
	profileType := model.ProfileType(profileTypeStr)
	if !model.IsValidProfileType(profileType) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile type"})
		return
	}

	profile, err := h.profileService.GetProfileByUserIDAndType(c.Request.Context(), userID, profileType)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	c.JSON(http.StatusOK, profile)
}

// GetDefaultProfile handles GET /profiles/user/:user_id/default
func (h *ProfileHandler) GetDefaultProfile(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	profile, err := h.profileService.GetDefaultProfile(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Default profile not found"})
		return
	}

	c.JSON(http.StatusOK, profile)
}

// GetActiveProfiles handles GET /profiles/user/:user_id/active
func (h *ProfileHandler) GetActiveProfiles(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	profiles, err := h.profileService.GetActiveProfiles(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"profiles": profiles,
		"count":    len(profiles),
	})
}

// UpdateProfileRequest represents the request to update a profile
type UpdateProfileRequest struct {
	IsActive  *bool `json:"is_active,omitempty"`
	IsDefault *bool `json:"is_default,omitempty"`
}

// UpdateProfile handles PUT /profiles/:id
func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile ID"})
		return
	}

	// Get existing profile
	profile, err := h.profileService.GetProfileByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	var req UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	// Update fields
	if req.IsActive != nil {
		profile.IsActive = *req.IsActive
	}
	if req.IsDefault != nil {
		profile.IsDefault = *req.IsDefault
		// If setting as default, update via service to handle unsetting other defaults
		if *req.IsDefault {
			if err := h.profileService.SetDefaultProfile(c.Request.Context(), profile.UserID, profile.ID); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}
	}

	if err := h.profileService.UpdateProfile(c.Request.Context(), profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully",
		"profile": profile,
	})
}

// SetDefaultProfile handles PATCH /profiles/:id/set-default
func (h *ProfileHandler) SetDefaultProfile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile ID"})
		return
	}

	// Get profile to get user ID
	profile, err := h.profileService.GetProfileByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	if err := h.profileService.SetDefaultProfile(c.Request.Context(), profile.UserID, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Default profile set successfully",
	})
}

// ActivateProfile handles PATCH /profiles/:id/activate
func (h *ProfileHandler) ActivateProfile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile ID"})
		return
	}

	if err := h.profileService.ActivateProfile(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile activated successfully",
	})
}

// DeactivateProfile handles PATCH /profiles/:id/deactivate
func (h *ProfileHandler) DeactivateProfile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile ID"})
		return
	}

	if err := h.profileService.DeactivateProfile(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile deactivated successfully",
	})
}

// DeleteProfile handles DELETE /profiles/:id
func (h *ProfileHandler) DeleteProfile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile ID"})
		return
	}

	if err := h.profileService.DeleteProfile(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile deleted successfully",
	})
}

// SyncProfiles handles POST /profiles/sync/:user_id
// Creates profiles for existing users who have role records but no profiles
func (h *ProfileHandler) SyncProfiles(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	profiles, err := h.profileService.SyncProfilesFromExistingRoles(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Profiles synced successfully",
		"profiles": profiles,
		"count":    len(profiles),
	})
}

// GetProfilesByType handles GET /profiles/type/:type
func (h *ProfileHandler) GetProfilesByType(c *gin.Context) {
	profileTypeStr := c.Param("type")
	profileType := model.ProfileType(profileTypeStr)
	if !model.IsValidProfileType(profileType) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile type"})
		return
	}

	// Parse pagination parameters
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	profiles, err := h.profileService.GetProfilesByType(c.Request.Context(), profileType, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"profiles": profiles,
		"count":    len(profiles),
		"limit":    limit,
		"offset":   offset,
	})
}
