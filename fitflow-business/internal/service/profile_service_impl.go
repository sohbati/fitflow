package service

import (
	"context"
	"errors"
	"fitflow-business/internal/model"
	"fitflow-business/internal/repository"
	"github.com/google/uuid"
)

// profileService implements ProfileService interface
type profileService struct {
	profileRepo  repository.ProfileRepository
	personRepo   repository.PersonRepository
	gymOwnerRepo repository.GymOwnerRepository
	trainerRepo  repository.TrainerRepository
	traineeRepo  repository.TraineeRepository
}

// NewProfileService creates a new profile service
func NewProfileService(
	profileRepo repository.ProfileRepository,
	personRepo repository.PersonRepository,
	gymOwnerRepo repository.GymOwnerRepository,
	trainerRepo repository.TrainerRepository,
	traineeRepo repository.TraineeRepository,
) ProfileService {
	return &profileService{
		profileRepo:  profileRepo,
		personRepo:   personRepo,
		gymOwnerRepo: gymOwnerRepo,
		trainerRepo:  trainerRepo,
		traineeRepo:  traineeRepo,
	}
}

// CreateProfile creates a new profile
func (s *profileService) CreateProfile(ctx context.Context, profile *model.Profile) error {
	// Validate profile
	if err := s.ValidateProfile(ctx, profile); err != nil {
		return err
	}

	// Check if person exists
	person, err := s.personRepo.GetPersonByID(ctx, profile.PersonID)
	if err != nil {
		return errors.New("person_not_found")
	}

	// Verify user_id matches
	if person.UserID != profile.UserID {
		return errors.New("user_id_does_not_match_person_user_id")
	}

	// Check if profile type already exists for this user (skip if creating from sync)
	existing, _ := s.profileRepo.GetProfileByUserIDAndType(ctx, profile.UserID, profile.Type)
	if existing != nil {
		return errors.New("profile_of_this_type_already_exists_for_this_user")
	}

	// If this is the first profile for the user, set it as default
	profiles, _ := s.profileRepo.GetProfilesByUserID(ctx, profile.UserID)
	if len(profiles) == 0 {
		profile.IsDefault = true
	}

	return s.profileRepo.CreateProfile(ctx, profile)
}

// GetProfileByID retrieves a profile by ID
func (s *profileService) GetProfileByID(ctx context.Context, id int64) (*model.Profile, error) {
	return s.profileRepo.GetProfileByID(ctx, id)
}

// GetProfilesByUserID retrieves all profiles for a user
func (s *profileService) GetProfilesByUserID(ctx context.Context, userID uuid.UUID) ([]*model.Profile, error) {
	return s.profileRepo.GetProfilesByUserID(ctx, userID)
}

// GetProfileByUserIDAndType retrieves a specific profile type for a user
func (s *profileService) GetProfileByUserIDAndType(ctx context.Context, userID uuid.UUID, profileType model.ProfileType) (*model.Profile, error) {
	return s.profileRepo.GetProfileByUserIDAndType(ctx, userID, profileType)
}

// GetDefaultProfile retrieves the default profile for a user
func (s *profileService) GetDefaultProfile(ctx context.Context, userID uuid.UUID) (*model.Profile, error) {
	return s.profileRepo.GetDefaultProfile(ctx, userID)
}

// UpdateProfile updates an existing profile
func (s *profileService) UpdateProfile(ctx context.Context, profile *model.Profile) error {
	// Validate profile
	if err := s.ValidateProfile(ctx, profile); err != nil {
		return err
	}

	return s.profileRepo.UpdateProfile(ctx, profile)
}

// DeleteProfile deletes a profile by ID
func (s *profileService) DeleteProfile(ctx context.Context, id int64) error {
	// Get the profile first to check if it's default
	profile, err := s.profileRepo.GetProfileByID(ctx, id)
	if err != nil {
		return err
	}

	// If deleting the default profile, set another profile as default
	if profile.IsDefault {
		profiles, _ := s.profileRepo.GetProfilesByUserID(ctx, profile.UserID)
		for _, p := range profiles {
			if p.ID != id {
				// Set the first non-deleted profile as default
				if err := s.profileRepo.SetDefaultProfile(ctx, profile.UserID, p.ID); err == nil {
					break
				}
			}
		}
	}

	return s.profileRepo.DeleteProfile(ctx, id)
}

// SetDefaultProfile sets a profile as the default for a user
func (s *profileService) SetDefaultProfile(ctx context.Context, userID uuid.UUID, profileID int64) error {
	// Verify the profile belongs to the user
	profile, err := s.profileRepo.GetProfileByID(ctx, profileID)
	if err != nil {
		return errors.New("profile_not_found")
	}

	if profile.UserID != userID {
		return errors.New("profile_does_not_belong_to_user")
	}

	return s.profileRepo.SetDefaultProfile(ctx, userID, profileID)
}

// GetActiveProfiles retrieves only active profiles for a user
func (s *profileService) GetActiveProfiles(ctx context.Context, userID uuid.UUID) ([]*model.Profile, error) {
	return s.profileRepo.GetActiveProfiles(ctx, userID)
}

// GetProfilesByType retrieves profiles by type with pagination
func (s *profileService) GetProfilesByType(ctx context.Context, profileType model.ProfileType, limit, offset int) ([]*model.Profile, error) {
	return s.profileRepo.GetProfilesByType(ctx, profileType, limit, offset)
}

// ActivateProfile activates a profile
func (s *profileService) ActivateProfile(ctx context.Context, id int64) error {
	return s.profileRepo.ActivateProfile(ctx, id)
}

// DeactivateProfile deactivates a profile
func (s *profileService) DeactivateProfile(ctx context.Context, id int64) error {
	return s.profileRepo.DeactivateProfile(ctx, id)
}

// CreateGymOwnerProfile creates a gym owner profile
func (s *profileService) CreateGymOwnerProfile(ctx context.Context, userID uuid.UUID, personID int64, gymOwnerID int64) (*model.Profile, error) {
	// Check if profile already exists
	existing, _ := s.profileRepo.GetProfileByUserIDAndType(ctx, userID, model.ProfileTypeGymOwner)
	if existing != nil {
		return existing, nil
	}

	profile := &model.Profile{
		UserID:     userID,
		Type:       model.ProfileTypeGymOwner,
		PersonID:   personID,
		GymOwnerID: &gymOwnerID,
		IsActive:   true,
	}

	if err := s.CreateProfile(ctx, profile); err != nil {
		return nil, err
	}

	return profile, nil
}

// CreateTrainerProfile creates a trainer profile
func (s *profileService) CreateTrainerProfile(ctx context.Context, userID uuid.UUID, personID int64, trainerID int64) (*model.Profile, error) {
	// Check if profile already exists
	existing, _ := s.profileRepo.GetProfileByUserIDAndType(ctx, userID, model.ProfileTypeTrainer)
	if existing != nil {
		return existing, nil
	}

	profile := &model.Profile{
		UserID:    userID,
		Type:      model.ProfileTypeTrainer,
		PersonID:  personID,
		TrainerID: &trainerID,
		IsActive:  true,
	}

	if err := s.CreateProfile(ctx, profile); err != nil {
		return nil, err
	}

	return profile, nil
}

// CreateTraineeProfile creates a trainee profile
func (s *profileService) CreateTraineeProfile(ctx context.Context, userID uuid.UUID, personID int64, traineeID int64) (*model.Profile, error) {
	// Check if profile already exists
	existing, _ := s.profileRepo.GetProfileByUserIDAndType(ctx, userID, model.ProfileTypeTrainee)
	if existing != nil {
		return existing, nil
	}

	profile := &model.Profile{
		UserID:    userID,
		Type:      model.ProfileTypeTrainee,
		PersonID:  personID,
		TraineeID: &traineeID,
		IsActive:  true,
	}

	if err := s.CreateProfile(ctx, profile); err != nil {
		return nil, err
	}

	return profile, nil
}

// ValidateProfile validates a profile
func (s *profileService) ValidateProfile(ctx context.Context, profile *model.Profile) error {
	// Validate profile type
	if !model.IsValidProfileType(profile.Type) {
		return errors.New("invalid_profile_type")
	}

	// Validate that the appropriate ID is set based on type
	switch profile.Type {
	case model.ProfileTypeGymOwner:
		if profile.GymOwnerID == nil {
			return errors.New("gym_owner_id_is_required_for_gym_owner_profile")
		}
		if profile.TrainerID != nil || profile.TraineeID != nil {
			return errors.New("only_gym_owner_id_should_be_set_for_gym_owner_profile")
		}
	case model.ProfileTypeTrainer:
		if profile.TrainerID == nil {
			return errors.New("trainer_id_is_required_for_trainer_profile")
		}
		if profile.GymOwnerID != nil || profile.TraineeID != nil {
			return errors.New("only_trainer_id_should_be_set_for_trainer_profile")
		}
	case model.ProfileTypeTrainee:
		if profile.TraineeID == nil {
			return errors.New("trainee_id_is_required_for_trainee_profile")
		}
		if profile.GymOwnerID != nil || profile.TrainerID != nil {
			return errors.New("only_trainee_id_should_be_set_for_trainee_profile")
		}
	}

	return nil
}

// SyncProfilesFromExistingRoles creates profiles for users who have role records but no profiles
// This is for backward compatibility with users registered before the profile system
func (s *profileService) SyncProfilesFromExistingRoles(ctx context.Context, userID uuid.UUID) ([]*model.Profile, error) {
	// Check if user already has profiles
	existingProfiles, _ := s.profileRepo.GetProfilesByUserID(ctx, userID)
	if len(existingProfiles) > 0 {
		// User already has profiles, return them
		return existingProfiles, nil
	}

	// Get person by user ID
	person, err := s.personRepo.GetPersonByUserID(ctx, userID)
	if err != nil {
		return nil, errors.New("person_not_found_for_user")
	}

	var createdProfiles []*model.Profile
	firstProfile := true

	// Check for gym owner
	gymOwner, err := s.gymOwnerRepo.GetGymOwnerByUserID(ctx, userID.String())
	if err == nil && gymOwner != nil {
		profile, err := s.CreateGymOwnerProfile(ctx, userID, person.ID, gymOwner.ID)
		if err == nil {
			if firstProfile {
				// Set first profile as default
				profile.IsDefault = true
				s.profileRepo.UpdateProfile(ctx, profile)
				firstProfile = false
			}
			createdProfiles = append(createdProfiles, profile)
		}
	}

	// Check for trainer (by person ID)
	trainer, err := s.trainerRepo.GetTrainerByPersonID(ctx, person.ID)
	if err == nil && trainer != nil {
		profile, err := s.CreateTrainerProfile(ctx, userID, person.ID, trainer.ID)
		if err == nil {
			if firstProfile {
				profile.IsDefault = true
				s.profileRepo.UpdateProfile(ctx, profile)
				firstProfile = false
			}
			createdProfiles = append(createdProfiles, profile)
		}
	}

	// Check for trainee (by person ID)
	trainee, err := s.traineeRepo.GetTraineeByPersonID(ctx, person.ID)
	if err == nil && trainee != nil {
		profile, err := s.CreateTraineeProfile(ctx, userID, person.ID, trainee.ID)
		if err == nil {
			if firstProfile {
				profile.IsDefault = true
				s.profileRepo.UpdateProfile(ctx, profile)
				firstProfile = false
			}
			createdProfiles = append(createdProfiles, profile)
		}
	}

	return createdProfiles, nil
}

