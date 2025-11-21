package impl

import (
	"context"
	"fitflow-business/internal/model"
	"fitflow-business/internal/repository"
	"gorm.io/gorm"
)

// gymLocationRepository implements GymLocationRepository interface
type gymLocationRepository struct {
	db *gorm.DB
}

// NewGymLocationRepository creates a new gym location repository
func NewGymLocationRepository(db *gorm.DB) repository.GymLocationRepository {
	return &gymLocationRepository{db: db}
}

// CreateGymLocation creates a new gym location
func (r *gymLocationRepository) CreateGymLocation(ctx context.Context, location *model.GymLocation) error {
	return r.db.WithContext(ctx).Create(location).Error
}

// GetGymLocations retrieves all locations for a gym
func (r *gymLocationRepository) GetGymLocations(ctx context.Context, gymID int64) ([]*model.GymLocation, error) {
	var locations []*model.GymLocation
	err := r.db.WithContext(ctx).Where("gym_id = ?", gymID).Find(&locations).Error
	return locations, err
}

// UpdateGymLocation updates an existing gym location
func (r *gymLocationRepository) UpdateGymLocation(ctx context.Context, location *model.GymLocation) error {
	return r.db.WithContext(ctx).Save(location).Error
}

// DeleteGymLocation deletes a gym location by ID
func (r *gymLocationRepository) DeleteGymLocation(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.GymLocation{}, id).Error
}
