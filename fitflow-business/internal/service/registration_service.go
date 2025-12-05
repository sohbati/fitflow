package service

import (
	"context"
	"errors"
	"fitflow-business/internal/model"
	"fitflow-business/internal/repository"
	"github.com/google/uuid"
)

// RegistrationService defines the interface for user registration operations
type RegistrationService interface {
	RegisterGymOwner(ctx context.Context, userID uuid.UUID, person *model.Person, gym *model.Gym, briefBio *string) (*model.GymOwner, error)
}

// registrationService implements RegistrationService interface
type registrationService struct {
	personService  PersonService
	gymService     GymService
	gymOwnerRepo   repository.GymOwnerRepository
	profileService ProfileService
}

// NewRegistrationService creates a new registration service
func NewRegistrationService(
	personService PersonService,
	gymService GymService,
	gymOwnerRepo repository.GymOwnerRepository,
	profileService ProfileService,
) RegistrationService {
	return &registrationService{
		personService:  personService,
		gymService:     gymService,
		gymOwnerRepo:   gymOwnerRepo,
		profileService: profileService,
	}
}

// RegisterGymOwner creates a person, gym, and gym_owner relationship
func (s *registrationService) RegisterGymOwner(ctx context.Context, userID uuid.UUID, person *model.Person, gym *model.Gym, briefBio *string) (*model.GymOwner, error) {
	// Validate inputs
	if userID == uuid.Nil {
		return nil, errors.New("user ID is required")
	}
	if person == nil {
		return nil, errors.New("person is required")
	}
	if gym == nil {
		return nil, errors.New("gym is required")
	}
	if person.FirstName == "" || person.LastName == "" {
		return nil, errors.New("first name and last name are required")
	}
	if gym.Name == "" {
		return nil, errors.New("gym name is required")
	}

	// Set user ID
	person.UserID = userID

	// Check if person already exists for this user
	existingPerson, err := s.personService.GetPersonByUserID(ctx, userID)
	if err == nil && existingPerson != nil {
		return nil, errors.New("person already exists for this user")
	}

	// Create person
	if err := s.personService.CreatePerson(ctx, person); err != nil {
		return nil, err
	}

	// Create gym
	if err := s.gymService.CreateGym(ctx, gym); err != nil {
		return nil, err
	}

	// Create gym owner relationship
	gymOwner := &model.GymOwner{
		PersonID: person.ID,
		GymID:    gym.ID,
		BriefBio: briefBio,
	}

	if err := s.gymOwnerRepo.CreateGymOwner(ctx, gymOwner); err != nil {
		return nil, err
	}

	// Create profile for gym owner
	_, err := s.profileService.CreateGymOwnerProfile(ctx, userID, person.ID, gymOwner.ID)
	if err != nil {
		// Log error but don't fail registration if profile creation fails
		// Profile can be created later via sync endpoint
		// In production, you might want to handle this differently
		// For now, we'll just continue
	}

	return gymOwner, nil
}

