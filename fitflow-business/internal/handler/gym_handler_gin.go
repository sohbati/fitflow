package handler

import (
	"fitflow-business/internal/model"
	"fitflow-business/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GymHandler handles HTTP requests for gym-related operations
type GymHandler struct {
	gymService service.GymService
}

// NewGymHandler creates a new gym handler
func NewGymHandler(gymService service.GymService) *GymHandler {
	return &GymHandler{
		gymService: gymService,
	}
}

// CreateGym handles POST /gyms
func (h *GymHandler) CreateGym(c *gin.Context) {
	var gym model.Gym
	if err := c.ShouldBindJSON(&gym); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := h.gymService.CreateGym(c.Request.Context(), &gym); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gym)
}

// GetGym handles GET /gyms/{id}
func (h *GymHandler) GetGym(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gym ID"})
		return
	}

	gym, err := h.gymService.GetGymByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gym)
}

// GetGyms handles GET /gyms
func (h *GymHandler) GetGyms(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	gyms, err := h.gymService.GetGyms(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gyms)
}

// UpdateGym handles PUT /gyms/{id}
func (h *GymHandler) UpdateGym(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gym ID"})
		return
	}

	var gym model.Gym
	if err := c.ShouldBindJSON(&gym); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	gym.ID = id
	if err := h.gymService.UpdateGym(c.Request.Context(), &gym); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gym)
}

// DeleteGym handles DELETE /gyms/{id}
func (h *GymHandler) DeleteGym(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gym ID"})
		return
	}

	if err := h.gymService.DeleteGym(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// SearchGyms handles GET /gyms/search
func (h *GymHandler) SearchGyms(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	gyms, err := h.gymService.SearchGyms(c.Request.Context(), query, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gyms)
}

// GetVerifiedGyms handles GET /gyms/verified
func (h *GymHandler) GetVerifiedGyms(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	gyms, err := h.gymService.GetVerifiedGyms(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gyms)
}

// CreateGymLocation handles POST /gyms/{id}/locations
func (h *GymHandler) CreateGymLocation(c *gin.Context) {
	gymID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gym ID"})
		return
	}

	var location model.GymLocation
	if err := c.ShouldBindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	location.GymID = gymID
	if err := h.gymService.CreateGymLocation(c.Request.Context(), &location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, location)
}

// GetGymLocations handles GET /gyms/{id}/locations
func (h *GymHandler) GetGymLocations(c *gin.Context) {
	gymID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gym ID"})
		return
	}

	locations, err := h.gymService.GetGymLocations(c.Request.Context(), gymID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, locations)
}

// CreateGymOwner handles POST /gyms/{id}/owners
func (h *GymHandler) CreateGymOwner(c *gin.Context) {
	gymID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gym ID"})
		return
	}

	var owner model.GymOwner
	if err := c.ShouldBindJSON(&owner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	owner.GymID = gymID
	if err := h.gymService.CreateGymOwner(c.Request.Context(), &owner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, owner)
}

// GetGymOwners handles GET /gyms/{id}/owners
func (h *GymHandler) GetGymOwners(c *gin.Context) {
	gymID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gym ID"})
		return
	}

	owners, err := h.gymService.GetGymOwners(c.Request.Context(), gymID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, owners)
}

// GetGymOwnerByUserID handles GET /gym-owners/user/:user_id
func (h *GymHandler) GetGymOwnerByUserID(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	gymOwner, err := h.gymService.GetGymOwnerByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gym owner not found"})
		return
	}

	c.JSON(http.StatusOK, gymOwner)
}

// CreateTrainer handles POST /trainers
func (h *GymHandler) CreateTrainer(c *gin.Context) {
	var trainer model.Trainer
	if err := c.ShouldBindJSON(&trainer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := h.gymService.CreateTrainer(c.Request.Context(), &trainer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, trainer)
}

// GetTrainer handles GET /trainers/{id}
func (h *GymHandler) GetTrainer(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid trainer ID"})
		return
	}

	trainer, err := h.gymService.GetTrainerByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, trainer)
}

// GetTrainerByUserID handles GET /trainers/user/:user_id
func (h *GymHandler) GetTrainerByUserID(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	trainer, err := h.gymService.GetTrainerByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Trainer not found"})
		return
	}

	c.JSON(http.StatusOK, trainer)
}

// GetTrainers handles GET /trainers
func (h *GymHandler) GetTrainers(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	trainers, err := h.gymService.GetTrainers(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, trainers)
}

// AddTrainerToGym handles POST /gyms/{id}/trainers/{trainer_id}
func (h *GymHandler) AddTrainerToGym(c *gin.Context) {
	gymID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gym ID"})
		return
	}

	trainerID, err := strconv.ParseInt(c.Param("trainer_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid trainer ID"})
		return
	}

	if err := h.gymService.AddTrainerToGym(c.Request.Context(), gymID, trainerID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// GetGymTrainers handles GET /gyms/{id}/trainers
func (h *GymHandler) GetGymTrainers(c *gin.Context) {
	gymID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gym ID"})
		return
	}

	trainers, err := h.gymService.GetGymTrainers(c.Request.Context(), gymID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, trainers)
}
