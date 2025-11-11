package impl

import (
	"context"
	"fitflow-business/internal/model"
	"gorm.io/gorm"
)

// gymRepository implements GymRepository interface
type gymRepository struct {
	db *gorm.DB
}

// NewGymRepository creates a new gym repository
func NewGymRepository(db *gorm.DB) GymRepository {
	return &gymRepository{db: db}
}

// CreateGym creates a new gym
func (r *gymRepository) CreateGym(ctx context.Context, gym *model.Gym) error {
	return r.db.WithContext(ctx).Create(gym).Error
}

// GetGymByID retrieves a gym by ID
func (r *gymRepository) GetGymByID(ctx context.Context, id int64) (*model.Gym, error) {
	var gym model.Gym
	err := r.db.WithContext(ctx).Preload("Locations").Preload("Owners").Preload("Trainers").First(&gym, id).Error
	if err != nil {
		return nil, err
	}
	return &gym, nil
}

// GetGyms retrieves all gyms with pagination
func (r *gymRepository) GetGyms(ctx context.Context, limit, offset int) ([]*model.Gym, error) {
	var gyms []*model.Gym
	err := r.db.WithContext(ctx).Preload("Locations").Preload("Owners").Preload("Trainers").
		Limit(limit).Offset(offset).Find(&gyms).Error
	return gyms, err
}

// UpdateGym updates an existing gym
func (r *gymRepository) UpdateGym(ctx context.Context, gym *model.Gym) error {
	return r.db.WithContext(ctx).Save(gym).Error
}

// DeleteGym deletes a gym by ID
func (r *gymRepository) DeleteGym(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.Gym{}, id).Error
}

// SearchGyms searches gyms by name or description
func (r *gymRepository) SearchGyms(ctx context.Context, query string, limit, offset int) ([]*model.Gym, error) {
	var gyms []*model.Gym
	err := r.db.WithContext(ctx).Preload("Locations").Preload("Owners").
		Where("name ILIKE ? OR description ILIKE ?", "%"+query+"%", "%"+query+"%").
		Limit(limit).Offset(offset).Find(&gyms).Error
	return gyms, err
}

// GetVerifiedGyms retrieves only verified gyms
func (r *gymRepository) GetVerifiedGyms(ctx context.Context, limit, offset int) ([]*model.Gym, error) {
	var gyms []*model.Gym
	err := r.db.WithContext(ctx).Preload("Locations").Preload("Owners").
		Where("is_verified = ?", true).
		Limit(limit).Offset(offset).Find(&gyms).Error
	return gyms, err
}