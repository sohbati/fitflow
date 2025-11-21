package impl

import (
	"context"
	"fitflow-business/internal/model"
	"fitflow-business/internal/repository"
	"gorm.io/gorm"
)

// gymOwnerRepository implements GymOwnerRepository interface
type gymOwnerRepository struct {
	db *gorm.DB
}

// NewGymOwnerRepository creates a new gym owner repository
func NewGymOwnerRepository(db *gorm.DB) repository.GymOwnerRepository {
	return &gymOwnerRepository{db: db}
}

// CreateGymOwner creates a new gym owner
func (r *gymOwnerRepository) CreateGymOwner(ctx context.Context, owner *model.GymOwner) error {
	return r.db.WithContext(ctx).Create(owner).Error
}

// GetGymOwners retrieves all owners for a gym
func (r *gymOwnerRepository) GetGymOwners(ctx context.Context, gymID int64) ([]*model.GymOwner, error) {
	var owners []*model.GymOwner
	err := r.db.WithContext(ctx).Where("gym_id = ?", gymID).Find(&owners).Error
	return owners, err
}

// UpdateGymOwner updates an existing gym owner
func (r *gymOwnerRepository) UpdateGymOwner(ctx context.Context, owner *model.GymOwner) error {
	return r.db.WithContext(ctx).Save(owner).Error
}

// DeleteGymOwner deletes a gym owner by ID
func (r *gymOwnerRepository) DeleteGymOwner(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.GymOwner{}, id).Error
}
