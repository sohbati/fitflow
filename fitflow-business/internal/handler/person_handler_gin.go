package handler

import (
	"fitflow-business/internal/model"
	"fitflow-business/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// PersonHandler handles HTTP requests for person-related operations
type PersonHandler struct {
	personService      service.PersonService
	registrationService service.RegistrationService
}

// NewPersonHandler creates a new person handler
func NewPersonHandler(personService service.PersonService, registrationService service.RegistrationService) *PersonHandler {
	return &PersonHandler{
		personService:      personService,
		registrationService: registrationService,
	}
}

// CheckPersonExists handles GET /persons/check/:user_id
// Checks if a person exists for the given user ID
func (h *PersonHandler) CheckPersonExists(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	person, err := h.personService.GetPersonByUserID(c.Request.Context(), userID)
	if err != nil {
		// Person doesn't exist
		c.JSON(http.StatusOK, gin.H{
			"exists": false,
			"person": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exists": true,
		"person": person,
	})
}

// RegisterGymOwnerRequest represents the request to register a gym owner
type RegisterGymOwnerRequest struct {
	UserID string `json:"user_id" binding:"required"` // User ID from IAM service

	// Person fields
	FirstName    string  `json:"first_name" binding:"required"`
	LastName     string  `json:"last_name" binding:"required"`
	Email        *string `json:"email"`
	PhoneNumber  *string `json:"phone_number"`
	DateOfBirth  *string `json:"date_of_birth"` // ISO date string
	Gender       *string `json:"gender"`        // "male", "female", "other"
	Address      *string `json:"address"`
	City         *string `json:"city"`
	Province     *string `json:"province"`
	Country      *string `json:"country"`
	PostalCode   *string `json:"postal_code"`
	ProfileImageURL *string `json:"profile_image_url"`

	// Gym fields
	GymName        string  `json:"gym_name" binding:"required"`
	GymDescription *string `json:"gym_description"`
	GymPhoneNumber *string `json:"gym_phone_number"`
	GymEmail       *string `json:"gym_email"`
	GymWebsiteURL  *string `json:"gym_website_url"`

	// Gym owner fields
	BriefBio *string `json:"brief_bio"`
}

// RegisterGymOwner handles POST /persons/register/gym-owner
// Creates a person, gym, and gym_owner relationship
func (h *PersonHandler) RegisterGymOwner(c *gin.Context) {
	var req RegisterGymOwnerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	// Parse user ID from request
	userUUID, err := uuid.Parse(req.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Parse date of birth if provided
	var dateOfBirth *time.Time
	if req.DateOfBirth != nil && *req.DateOfBirth != "" {
		parsed, err := time.Parse("2006-01-02", *req.DateOfBirth)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date of birth format. Use YYYY-MM-DD"})
			return
		}
		dateOfBirth = &parsed
	}

	// Parse gender if provided
	var gender *model.Gender
	if req.Gender != nil && *req.Gender != "" {
		g := model.Gender(*req.Gender)
		if g != model.GenderMale && g != model.GenderFemale && g != model.GenderOther {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gender. Must be 'male', 'female', or 'other'"})
			return
		}
		gender = &g
	}

	// Create person model
	person := &model.Person{
		UserID:          userUUID,
		FirstName:       req.FirstName,
		LastName:        req.LastName,
		Email:           req.Email,
		PhoneNumber:     req.PhoneNumber,
		DateOfBirth:     dateOfBirth,
		Gender:          gender,
		ProfileImageURL: req.ProfileImageURL,
		Address:         req.Address,
		City:            req.City,
		Province:        req.Province,
		Country:         req.Country,
		PostalCode:      req.PostalCode,
		IsActive:        true,
	}

	// Create gym model
	gym := &model.Gym{
		Name:        req.GymName,
		Description: req.GymDescription,
		PhoneNumber: req.GymPhoneNumber,
		Email:       req.GymEmail,
		WebsiteURL:  req.GymWebsiteURL,
		IsVerified:  false, // New gyms start as unverified
	}

	// Register gym owner
	gymOwner, err := h.registrationService.RegisterGymOwner(c.Request.Context(), userUUID, person, gym, req.BriefBio)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":   "Gym owner registered successfully",
		"gym_owner": gymOwner,
	})
}

