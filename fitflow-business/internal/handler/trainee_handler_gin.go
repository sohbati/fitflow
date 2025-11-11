package handler

import (
	"net/http"
	"strconv"
	"fitflow-business/internal/model"
	"fitflow-business/internal/service"
	"github.com/gin-gonic/gin"
)

// TraineeHandler handles HTTP requests for trainee operations
type TraineeHandler struct {
	traineeService service.TraineeService
}

// NewTraineeHandler creates a new trainee handler
func NewTraineeHandler(traineeService service.TraineeService) *TraineeHandler {
	return &TraineeHandler{
		traineeService: traineeService,
	}
}

// CreateTrainee handles POST /trainees
func (h *TraineeHandler) CreateTrainee(c *gin.Context) {
	var trainee model.Trainee
	if err := c.ShouldBindJSON(&trainee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.traineeService.CreateTrainee(c.Request.Context(), &trainee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, trainee)
}

// GetTrainee handles GET /trainees/:id
func (h *TraineeHandler) GetTrainee(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid trainee ID"})
		return
	}

	trainee, err := h.traineeService.GetTraineeByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, trainee)
}

// GetTrainees handles GET /trainees
func (h *TraineeHandler) GetTrainees(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit parameter"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset parameter"})
		return
	}

	trainees, err := h.traineeService.GetTrainees(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"trainees": trainees})
}

// UpdateTrainee handles PUT /trainees/:id
func (h *TraineeHandler) UpdateTrainee(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid trainee ID"})
		return
	}

	var trainee model.Trainee
	if err := c.ShouldBindJSON(&trainee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	trainee.ID = id
	if err := h.traineeService.UpdateTrainee(c.Request.Context(), &trainee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, trainee)
}

// DeleteTrainee handles DELETE /trainees/:id
func (h *TraineeHandler) DeleteTrainee(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid trainee ID"})
		return
	}

	if err := h.traineeService.DeleteTrainee(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// SearchTrainees handles GET /trainees/search
func (h *TraineeHandler) SearchTrainees(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "search query is required"})
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit parameter"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset parameter"})
		return
	}

	trainees, err := h.traineeService.SearchTrainees(c.Request.Context(), query, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"trainees": trainees})
}

// GetTraineesByEmail handles GET /trainees/email/:email
func (h *TraineeHandler) GetTraineesByEmail(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
		return
	}

	trainee, err := h.traineeService.GetTraineesByEmail(c.Request.Context(), email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, trainee)
}

// GetTraineesByPhone handles GET /trainees/phone/:phone
func (h *TraineeHandler) GetTraineesByPhone(c *gin.Context) {
	phone := c.Param("phone")
	if phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "phone number is required"})
		return
	}

	trainee, err := h.traineeService.GetTraineesByPhone(c.Request.Context(), phone)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, trainee)
}

// GetActiveTrainees handles GET /trainees/active
func (h *TraineeHandler) GetActiveTrainees(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit parameter"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset parameter"})
		return
	}

	trainees, err := h.traineeService.GetActiveTrainees(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"trainees": trainees})
}

// GetTraineesByMembershipType handles GET /trainees/membership/:type
func (h *TraineeHandler) GetTraineesByMembershipType(c *gin.Context) {
	membershipType := c.Param("type")
	if membershipType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "membership type is required"})
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit parameter"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset parameter"})
		return
	}

	trainees, err := h.traineeService.GetTraineesByMembershipType(c.Request.Context(), model.MembershipType(membershipType), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"trainees": trainees})
}

// GetTraineesByFitnessLevel handles GET /trainees/fitness/:level
func (h *TraineeHandler) GetTraineesByFitnessLevel(c *gin.Context) {
	fitnessLevel := c.Param("level")
	if fitnessLevel == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "fitness level is required"})
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit parameter"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset parameter"})
		return
	}

	trainees, err := h.traineeService.GetTraineesByFitnessLevel(c.Request.Context(), model.FitnessLevel(fitnessLevel), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"trainees": trainees})
}

// GetExpiringMemberships handles GET /trainees/expiring
func (h *TraineeHandler) GetExpiringMemberships(c *gin.Context) {
	daysStr := c.DefaultQuery("days", "30")
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	days, err := strconv.Atoi(daysStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid days parameter"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit parameter"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset parameter"})
		return
	}

	trainees, err := h.traineeService.GetExpiringMemberships(c.Request.Context(), days, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"trainees": trainees})
}

// GetExpiredMemberships handles GET /trainees/expired
func (h *TraineeHandler) GetExpiredMemberships(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit parameter"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset parameter"})
		return
	}

	trainees, err := h.traineeService.GetExpiredMemberships(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"trainees": trainees})
}

// UpdateMembershipStatus handles PATCH /trainees/:id/status
func (h *TraineeHandler) UpdateMembershipStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid trainee ID"})
		return
	}

	var request struct {
		IsActive bool `json:"is_active"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.traineeService.UpdateMembershipStatus(c.Request.Context(), id, request.IsActive); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "membership status updated successfully"})
}

// CalculateAge handles GET /trainees/:id/age
func (h *TraineeHandler) CalculateAge(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid trainee ID"})
		return
	}

	trainee, err := h.traineeService.GetTraineeByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	age, err := h.traineeService.CalculateAge(c.Request.Context(), trainee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"age": age})
}

// CalculateBMI handles GET /trainees/:id/bmi
func (h *TraineeHandler) CalculateBMI(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid trainee ID"})
		return
	}

	trainee, err := h.traineeService.GetTraineeByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	bmi, err := h.traineeService.CalculateBMI(c.Request.Context(), trainee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"bmi": bmi})
}

// ValidateMembership handles GET /trainees/:id/membership/validate
func (h *TraineeHandler) ValidateMembership(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid trainee ID"})
		return
	}

	trainee, err := h.traineeService.GetTraineeByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	isValid, err := h.traineeService.ValidateMembership(c.Request.Context(), trainee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"is_valid": isValid})
}
