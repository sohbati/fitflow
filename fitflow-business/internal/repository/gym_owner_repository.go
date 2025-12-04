package repository

import (
	"context"
	"fitflow-business/internal/model"
)

// GymOwnerRepository defines the interface for gym owner database operations
type GymOwnerRepository interface {
	CreateGymOwner(ctx context.Context, owner *model.GymOwner) error
	GetGymOwners(ctx context.Context, gymID int64) ([]*model.GymOwner, error)
	GetGymOwnerByPersonID(ctx context.Context, personID int64) (*model.GymOwner, error)
	GetGymOwnerByUserID(ctx context.Context, userID string) (*model.GymOwner, error)
	UpdateGymOwner(ctx context.Context, owner *model.GymOwner) error
	DeleteGymOwner(ctx context.Context, id int64) error
}
