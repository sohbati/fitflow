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
	RegisterTrainer(ctx context.Context, userID uuid.UUID, person *model.Person) (*model.Trainer, error)
	RegisterTrainee(ctx context.Context, userID uuid.UUID, person *model.Person, trainee *model.Trainee) (*model.Trainee, error)
}

// registrationService implements RegistrationService interface
type registrationService struct {
	personService  PersonService
	gymService     GymService
	gymOwnerRepo   repository.GymOwnerRepository
	trainerRepo    repository.TrainerRepository
	traineeRepo    repository.TraineeRepository
	profileService ProfileService
}

// NewRegistrationService creates a new registration service
func NewRegistrationService(
	personService PersonService,
	gymService GymService,
	gymOwnerRepo repository.GymOwnerRepository,
	trainerRepo repository.TrainerRepository,
	traineeRepo repository.TraineeRepository,
	profileService ProfileService,
) RegistrationService {
	return &registrationService{
		personService:  personService,
		gymService:     gymService,
		gymOwnerRepo:   gymOwnerRepo,
		trainerRepo:    trainerRepo,
		traineeRepo:    traineeRepo,
		profileService: profileService,
	}
}

// RegisterGymOwner creates a person, gym, and gym_owner relationship
func (s *registrationService) RegisterGymOwner(ctx context.Context, userID uuid.UUID, person *model.Person, gym *model.Gym, briefBio *string) (*model.GymOwner, error) {
	// Validate inputs
	if userID == uuid.Nil {
		return nil, errors.New("user_id_is_required")
	}
	if person == nil {
		return nil, errors.New("person_is_required")
	}
	if gym == nil {
		return nil, errors.New("gym_is_required")
	}
	if person.FirstName == "" || person.LastName == "" {
		return nil, errors.New("first_name_and_last_name_are_required")
	}
	if gym.Name == "" {
		return nil, errors.New("gym_name_is_required")
	}

	// Set user ID
	person.UserID = userID

	// Check if person already exists for this user
	existingPerson, err := s.personService.GetPersonByUserID(ctx, userID)
	if err == nil && existingPerson != nil {
		return nil, errors.New("person_already_exists_for_this_user")
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
	_, err = s.profileService.CreateGymOwnerProfile(ctx, userID, person.ID, gymOwner.ID)
	if err != nil {
		// Log error but don't fail registration if profile creation fails
		// Profile can be created later via sync endpoint
		// In production, you might want to handle this differently
		// For now, we'll just continue
	}

	return gymOwner, nil
}

// RegisterTrainer creates a person and trainer relationship
func (s *registrationService) RegisterTrainer(ctx context.Context, userID uuid.UUID, person *model.Person) (*model.Trainer, error) {
	// Validate inputs
	if userID == uuid.Nil {
		return nil, errors.New("user_id_is_required")
	}
	if person == nil {
		return nil, errors.New("person_is_required")
	}
	if person.FirstName == "" || person.LastName == "" {
		return nil, errors.New("first_name_and_last_name_are_required")
	}

	// Set user ID
	person.UserID = userID

	// Get or create person
	existingPerson, err := s.personService.GetPersonByUserID(ctx, userID)
	if err == nil && existingPerson != nil {
		// Person exists, use it
		person = existingPerson
	} else {
		// Create person
		if err := s.personService.CreatePerson(ctx, person); err != nil {
			return nil, err
		}
	}

	// Check if trainer already exists for this person
	existingTrainer, err := s.trainerRepo.GetTrainerByPersonID(ctx, person.ID)
	if err == nil && existingTrainer != nil {
		return nil, errors.New("trainer_already_exists_for_this_person")
	}

	// Create trainer
	trainer := &model.Trainer{
		PersonID:     person.ID,
		IsRegistered: false, // Will be set to true when they register with a gym
	}

	if err := s.trainerRepo.CreateTrainer(ctx, trainer); err != nil {
		return nil, err
	}

	// Create profile for trainer
	_, err = s.profileService.CreateTrainerProfile(ctx, userID, person.ID, trainer.ID)
	if err != nil {
		// Log error but don't fail registration if profile creation fails
	}

	return trainer, nil
}

// RegisterTrainee creates a person and trainee relationship
func (s *registrationService) RegisterTrainee(ctx context.Context, userID uuid.UUID, person *model.Person, trainee *model.Trainee) (*model.Trainee, error) {
	// Validate inputs
	if userID == uuid.Nil {
		return nil, errors.New("user_id_is_required")
	}
	if person == nil {
		return nil, errors.New("person_is_required")
	}
	if person.FirstName == "" || person.LastName == "" {
		return nil, errors.New("first_name_and_last_name_are_required")
	}
	if trainee == nil {
		return nil, errors.New("trainee_is_required")
	}

	// Set user ID
	person.UserID = userID

	// Get or create person
	existingPerson, err := s.personService.GetPersonByUserID(ctx, userID)
	if err == nil && existingPerson != nil {
		// Person exists, use it
		person = existingPerson
	} else {
		// Create person
		if err := s.personService.CreatePerson(ctx, person); err != nil {
			return nil, err
		}
	}

	// Check if trainee already exists for this person
	existingTrainee, err := s.traineeRepo.GetTraineeByPersonID(ctx, person.ID)
	if err == nil && existingTrainee != nil {
		return nil, errors.New("trainee_already_exists_for_this_person")
	}

	// Set person ID for trainee
	trainee.PersonID = person.ID

	// Create trainee
	if err := s.traineeRepo.CreateTrainee(ctx, trainee); err != nil {
		return nil, err
	}

	// Create profile for trainee
	_, err = s.profileService.CreateTraineeProfile(ctx, userID, person.ID, trainee.ID)
	if err != nil {
		// Log error but don't fail registration if profile creation fails
	}

	return trainee, nil
}
