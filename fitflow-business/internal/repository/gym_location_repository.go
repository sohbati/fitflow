package repository

import (
	"context"
	"fitflow-business/internal/model"
)

// GymLocationRepository defines the interface for gym location database operations
type GymLocationRepository interface {
	CreateGymLocation(ctx context.Context, location *model.GymLocation) error
	GetGymLocations(ctx context.Context, gymID int64) ([]*model.GymLocation, error)
	UpdateGymLocation(ctx context.Context, location *model.GymLocation) error
	DeleteGymLocation(ctx context.Context, id int64) error
}
