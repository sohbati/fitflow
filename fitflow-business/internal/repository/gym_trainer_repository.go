package repository

import (
	"context"
	"fitflow-business/internal/model"
)

// GymTrainerRepository defines the interface for gym-trainer relationship database operations
type GymTrainerRepository interface {
	AddTrainerToGym(ctx context.Context, gymID, trainerID int64) error
	RemoveTrainerFromGym(ctx context.Context, gymID, trainerID int64) error
	GetGymTrainers(ctx context.Context, gymID int64) ([]*model.Trainer, error)
	GetTrainerGyms(ctx context.Context, trainerID int64) ([]*model.Gym, error)
}
