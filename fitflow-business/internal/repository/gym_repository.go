package repository

import (
	"context"
	"fitflow-business/internal/model"
)

// GymRepository defines the interface for gym-related database operations
type GymRepository interface {
	// Gym operations
	CreateGym(ctx context.Context, gym *model.Gym) error
	GetGymByID(ctx context.Context, id int64) (*model.Gym, error)
	GetGyms(ctx context.Context, limit, offset int) ([]*model.Gym, error)
	UpdateGym(ctx context.Context, gym *model.Gym) error
	DeleteGym(ctx context.Context, id int64) error
	SearchGyms(ctx context.Context, query string, limit, offset int) ([]*model.Gym, error)
	GetVerifiedGyms(ctx context.Context, limit, offset int) ([]*model.Gym, error)
}