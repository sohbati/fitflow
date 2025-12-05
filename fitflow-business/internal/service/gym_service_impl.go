package service

import (
	"context"
	"errors"
	"fitflow-business/internal/model"
	"fitflow-business/internal/repository"
)

// gymService implements GymService interface
type gymService struct {
	gymRepo         repository.GymRepository
	gymLocationRepo repository.GymLocationRepository
	gymOwnerRepo    repository.GymOwnerRepository
	trainerRepo     repository.TrainerRepository
	gymTrainerRepo  repository.GymTrainerRepository
}

// NewGymService creates a new gym service
func NewGymService(
	gymRepo repository.GymRepository,
	gymLocationRepo repository.GymLocationRepository,
	gymOwnerRepo repository.GymOwnerRepository,
	trainerRepo repository.TrainerRepository,
	gymTrainerRepo repository.GymTrainerRepository,
) GymService {
	return &gymService{
		gymRepo:         gymRepo,
		gymLocationRepo: gymLocationRepo,
		gymOwnerRepo:    gymOwnerRepo,
		trainerRepo:     trainerRepo,
		gymTrainerRepo:  gymTrainerRepo,
	}
}

// CreateGym creates a new gym with validation
func (s *gymService) CreateGym(ctx context.Context, gym *model.Gym) error {
	// Validate required fields
	if gym.Name == "" {
		return errors.New("gym name is required")
	}

	// Set default values
	if gym.Facilities == nil {
		gym.Facilities = make(model.JSONB)
	}
	if gym.Images == nil {
		gym.Images = make(model.Images, 0)
	}

	return s.gymRepo.CreateGym(ctx, gym)
}

// GetGymByID retrieves a gym by ID
func (s *gymService) GetGymByID(ctx context.Context, id int64) (*model.Gym, error) {
	if id <= 0 {
		return nil, errors.New("invalid gym ID")
	}
	return s.gymRepo.GetGymByID(ctx, id)
}

// GetGyms retrieves all gyms with pagination
func (s *gymService) GetGyms(ctx context.Context, limit, offset int) ([]*model.Gym, error) {
	if limit <= 0 {
		limit = 10 // Default limit
	}
	if limit > 100 {
		limit = 100 // Max limit
	}
	if offset < 0 {
		offset = 0
	}

	return s.gymRepo.GetGyms(ctx, limit, offset)
}

// UpdateGym updates an existing gym
func (s *gymService) UpdateGym(ctx context.Context, gym *model.Gym) error {
	if gym.ID <= 0 {
		return errors.New("invalid gym ID")
	}
	if gym.Name == "" {
		return errors.New("gym name is required")
	}

	return s.gymRepo.UpdateGym(ctx, gym)
}

// DeleteGym deletes a gym by ID
func (s *gymService) DeleteGym(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid gym ID")
	}
	return s.gymRepo.DeleteGym(ctx, id)
}

// SearchGyms searches gyms by name or description
func (s *gymService) SearchGyms(ctx context.Context, query string, limit, offset int) ([]*model.Gym, error) {
	if query == "" {
		return nil, errors.New("search query cannot be empty")
	}
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	return s.gymRepo.SearchGyms(ctx, query, limit, offset)
}

// GetVerifiedGyms retrieves only verified gyms
func (s *gymService) GetVerifiedGyms(ctx context.Context, limit, offset int) ([]*model.Gym, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	return s.gymRepo.GetVerifiedGyms(ctx, limit, offset)
}

// CreateGymLocation creates a new gym location
func (s *gymService) CreateGymLocation(ctx context.Context, location *model.GymLocation) error {
	if location.GymID <= 0 {
		return errors.New("invalid gym ID")
	}
	if location.LocationType == "" {
		return errors.New("location type is required")
	}

	return s.gymLocationRepo.CreateGymLocation(ctx, location)
}

// GetGymLocations retrieves all locations for a gym
func (s *gymService) GetGymLocations(ctx context.Context, gymID int64) ([]*model.GymLocation, error) {
	if gymID <= 0 {
		return nil, errors.New("invalid gym ID")
	}
	return s.gymLocationRepo.GetGymLocations(ctx, gymID)
}

// UpdateGymLocation updates an existing gym location
func (s *gymService) UpdateGymLocation(ctx context.Context, location *model.GymLocation) error {
	if location.ID <= 0 {
		return errors.New("invalid location ID")
	}
	if location.LocationType == "" {
		return errors.New("location type is required")
	}

	return s.gymLocationRepo.UpdateGymLocation(ctx, location)
}

// DeleteGymLocation deletes a gym location by ID
func (s *gymService) DeleteGymLocation(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid location ID")
	}
	return s.gymLocationRepo.DeleteGymLocation(ctx, id)
}

// CreateGymOwner creates a new gym owner
func (s *gymService) CreateGymOwner(ctx context.Context, owner *model.GymOwner) error {
	if owner.GymID <= 0 {
		return errors.New("invalid gym ID")
	}
	if owner.PersonID <= 0 && owner.Person.ID <= 0 {
		return errors.New("owner person reference is required")
	}

	return s.gymOwnerRepo.CreateGymOwner(ctx, owner)
}

// GetGymOwners retrieves all owners for a gym
func (s *gymService) GetGymOwners(ctx context.Context, gymID int64) ([]*model.GymOwner, error) {
	if gymID <= 0 {
		return nil, errors.New("invalid gym ID")
	}
	return s.gymOwnerRepo.GetGymOwners(ctx, gymID)
}

// GetGymOwnerByUserID retrieves a gym owner by user ID
func (s *gymService) GetGymOwnerByUserID(ctx context.Context, userID string) (*model.GymOwner, error) {
	if userID == "" {
		return nil, errors.New("user ID is required")
	}
	return s.gymOwnerRepo.GetGymOwnerByUserID(ctx, userID)
}

// UpdateGymOwner updates an existing gym owner
func (s *gymService) UpdateGymOwner(ctx context.Context, owner *model.GymOwner) error {
	if owner.ID <= 0 {
		return errors.New("invalid owner ID")
	}
	if owner.PersonID <= 0 && owner.Person.ID <= 0 {
		return errors.New("owner person reference is required")
	}

	return s.gymOwnerRepo.UpdateGymOwner(ctx, owner)
}

// DeleteGymOwner deletes a gym owner by ID
func (s *gymService) DeleteGymOwner(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid owner ID")
	}
	return s.gymOwnerRepo.DeleteGymOwner(ctx, id)
}

// CreateTrainer creates a new trainer
func (s *gymService) CreateTrainer(ctx context.Context, trainer *model.Trainer) error {
	if trainer.PersonID <= 0 && trainer.Person.ID <= 0 {
		return errors.New("trainer person reference is required")
	}

	return s.trainerRepo.CreateTrainer(ctx, trainer)
}

// GetTrainerByID retrieves a trainer by ID
func (s *gymService) GetTrainerByID(ctx context.Context, id int64) (*model.Trainer, error) {
	if id <= 0 {
		return nil, errors.New("invalid trainer ID")
	}
	return s.trainerRepo.GetTrainerByID(ctx, id)
}

// GetTrainerByUserID retrieves a trainer by user ID
func (s *gymService) GetTrainerByUserID(ctx context.Context, userID string) (*model.Trainer, error) {
	if userID == "" {
		return nil, errors.New("user ID is required")
	}
	return s.trainerRepo.GetTrainerByUserID(ctx, userID)
}

// GetTrainers retrieves all trainers with pagination
func (s *gymService) GetTrainers(ctx context.Context, limit, offset int) ([]*model.Trainer, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	return s.trainerRepo.GetTrainers(ctx, limit, offset)
}

// UpdateTrainer updates an existing trainer
func (s *gymService) UpdateTrainer(ctx context.Context, trainer *model.Trainer) error {
	if trainer.ID <= 0 {
		return errors.New("invalid trainer ID")
	}
	if trainer.PersonID <= 0 && trainer.Person.ID <= 0 {
		return errors.New("trainer person reference is required")
	}

	return s.trainerRepo.UpdateTrainer(ctx, trainer)
}

// DeleteTrainer deletes a trainer by ID
func (s *gymService) DeleteTrainer(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid trainer ID")
	}
	return s.trainerRepo.DeleteTrainer(ctx, id)
}

// GetRegisteredTrainers retrieves only registered trainers
func (s *gymService) GetRegisteredTrainers(ctx context.Context, limit, offset int) ([]*model.Trainer, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	return s.trainerRepo.GetRegisteredTrainers(ctx, limit, offset)
}

// AddTrainerToGym adds a trainer to a gym
func (s *gymService) AddTrainerToGym(ctx context.Context, gymID, trainerID int64) error {
	if gymID <= 0 {
		return errors.New("invalid gym ID")
	}
	if trainerID <= 0 {
		return errors.New("invalid trainer ID")
	}

	return s.gymTrainerRepo.AddTrainerToGym(ctx, gymID, trainerID)
}

// RemoveTrainerFromGym removes a trainer from a gym
func (s *gymService) RemoveTrainerFromGym(ctx context.Context, gymID, trainerID int64) error {
	if gymID <= 0 {
		return errors.New("invalid gym ID")
	}
	if trainerID <= 0 {
		return errors.New("invalid trainer ID")
	}

	return s.gymTrainerRepo.RemoveTrainerFromGym(ctx, gymID, trainerID)
}

// GetGymTrainers retrieves all trainers for a gym
func (s *gymService) GetGymTrainers(ctx context.Context, gymID int64) ([]*model.Trainer, error) {
	if gymID <= 0 {
		return nil, errors.New("invalid gym ID")
	}
	return s.gymTrainerRepo.GetGymTrainers(ctx, gymID)
}

// GetTrainerGyms retrieves all gyms for a trainer
func (s *gymService) GetTrainerGyms(ctx context.Context, trainerID int64) ([]*model.Gym, error) {
	if trainerID <= 0 {
		return nil, errors.New("invalid trainer ID")
	}
	return s.gymTrainerRepo.GetTrainerGyms(ctx, trainerID)
}
