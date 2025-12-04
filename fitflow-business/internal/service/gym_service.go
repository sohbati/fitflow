package service

import (
	"context"
	"fitflow-business/internal/model"
)

// GymService defines the interface for gym-related business logic
type GymService interface {
	// Gym operations
	CreateGym(ctx context.Context, gym *model.Gym) error
	GetGymByID(ctx context.Context, id int64) (*model.Gym, error)
	GetGyms(ctx context.Context, limit, offset int) ([]*model.Gym, error)
	UpdateGym(ctx context.Context, gym *model.Gym) error
	DeleteGym(ctx context.Context, id int64) error
	SearchGyms(ctx context.Context, query string, limit, offset int) ([]*model.Gym, error)
	GetVerifiedGyms(ctx context.Context, limit, offset int) ([]*model.Gym, error)

	// Gym location operations
	CreateGymLocation(ctx context.Context, location *model.GymLocation) error
	GetGymLocations(ctx context.Context, gymID int64) ([]*model.GymLocation, error)
	UpdateGymLocation(ctx context.Context, location *model.GymLocation) error
	DeleteGymLocation(ctx context.Context, id int64) error

	// Gym owner operations
	CreateGymOwner(ctx context.Context, owner *model.GymOwner) error
	GetGymOwners(ctx context.Context, gymID int64) ([]*model.GymOwner, error)
	GetGymOwnerByUserID(ctx context.Context, userID string) (*model.GymOwner, error)
	UpdateGymOwner(ctx context.Context, owner *model.GymOwner) error
	DeleteGymOwner(ctx context.Context, id int64) error

	// Trainer operations
	CreateTrainer(ctx context.Context, trainer *model.Trainer) error
	GetTrainerByID(ctx context.Context, id int64) (*model.Trainer, error)
	GetTrainers(ctx context.Context, limit, offset int) ([]*model.Trainer, error)
	UpdateTrainer(ctx context.Context, trainer *model.Trainer) error
	DeleteTrainer(ctx context.Context, id int64) error
	GetRegisteredTrainers(ctx context.Context, limit, offset int) ([]*model.Trainer, error)

	// Gym-Trainer relationship operations
	AddTrainerToGym(ctx context.Context, gymID, trainerID int64) error
	RemoveTrainerFromGym(ctx context.Context, gymID, trainerID int64) error
	GetGymTrainers(ctx context.Context, gymID int64) ([]*model.Trainer, error)
	GetTrainerGyms(ctx context.Context, trainerID int64) ([]*model.Gym, error)
}
