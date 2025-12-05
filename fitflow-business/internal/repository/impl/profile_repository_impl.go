package impl

import (
	"context"
	"fitflow-business/internal/model"
	"fitflow-business/internal/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// profileRepository implements ProfileRepository interface
type profileRepository struct {
	db *gorm.DB
}

// NewProfileRepository creates a new profile repository
func NewProfileRepository(db *gorm.DB) repository.ProfileRepository {
	return &profileRepository{db: db}
}

// CreateProfile creates a new profile
func (r *profileRepository) CreateProfile(ctx context.Context, profile *model.Profile) error {
	return r.db.WithContext(ctx).Create(profile).Error
}

// GetProfileByID retrieves a profile by ID
func (r *profileRepository) GetProfileByID(ctx context.Context, id int64) (*model.Profile, error) {
	var profile model.Profile
	err := r.db.WithContext(ctx).Preload("Person").First(&profile, id).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

// GetProfilesByUserID retrieves all profiles for a user
func (r *profileRepository) GetProfilesByUserID(ctx context.Context, userID uuid.UUID) ([]*model.Profile, error) {
	var profiles []*model.Profile
	err := r.db.WithContext(ctx).
		Preload("Person").
		Where("user_id = ?", userID).
		Order("is_default DESC, created_at ASC").
		Find(&profiles).Error
	return profiles, err
}

// GetProfileByUserIDAndType retrieves a specific profile type for a user
func (r *profileRepository) GetProfileByUserIDAndType(ctx context.Context, userID uuid.UUID, profileType model.ProfileType) (*model.Profile, error) {
	var profile model.Profile
	err := r.db.WithContext(ctx).
		Preload("Person").
		Where("user_id = ? AND type = ?", userID, profileType).
		First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

// GetDefaultProfile retrieves the default profile for a user
func (r *profileRepository) GetDefaultProfile(ctx context.Context, userID uuid.UUID) (*model.Profile, error) {
	var profile model.Profile
	err := r.db.WithContext(ctx).
		Preload("Person").
		Where("user_id = ? AND is_default = ?", userID, true).
		First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

// UpdateProfile updates an existing profile
func (r *profileRepository) UpdateProfile(ctx context.Context, profile *model.Profile) error {
	return r.db.WithContext(ctx).Save(profile).Error
}

// DeleteProfile deletes a profile by ID
func (r *profileRepository) DeleteProfile(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.Profile{}, id).Error
}

// SetDefaultProfile sets a profile as the default for a user
// This will unset any other default profiles for the same user
func (r *profileRepository) SetDefaultProfile(ctx context.Context, userID uuid.UUID, profileID int64) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// First, unset all default profiles for this user
		if err := tx.Model(&model.Profile{}).
			Where("user_id = ?", userID).
			Update("is_default", false).Error; err != nil {
			return err
		}

		// Then set the specified profile as default
		if err := tx.Model(&model.Profile{}).
			Where("id = ? AND user_id = ?", profileID, userID).
			Update("is_default", true).Error; err != nil {
			return err
		}

		return nil
	})
}

// GetActiveProfiles retrieves only active profiles for a user
func (r *profileRepository) GetActiveProfiles(ctx context.Context, userID uuid.UUID) ([]*model.Profile, error) {
	var profiles []*model.Profile
	err := r.db.WithContext(ctx).
		Preload("Person").
		Where("user_id = ? AND is_active = ?", userID, true).
		Order("is_default DESC, created_at ASC").
		Find(&profiles).Error
	return profiles, err
}

// GetProfilesByType retrieves profiles by type with pagination
func (r *profileRepository) GetProfilesByType(ctx context.Context, profileType model.ProfileType, limit, offset int) ([]*model.Profile, error) {
	var profiles []*model.Profile
	err := r.db.WithContext(ctx).
		Preload("Person").
		Where("type = ?", profileType).
		Limit(limit).Offset(offset).
		Find(&profiles).Error
	return profiles, err
}

// ActivateProfile activates a profile
func (r *profileRepository) ActivateProfile(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).
		Model(&model.Profile{}).
		Where("id = ?", id).
		Update("is_active", true).Error
}

// DeactivateProfile deactivates a profile
func (r *profileRepository) DeactivateProfile(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).
		Model(&model.Profile{}).
		Where("id = ?", id).
		Update("is_active", false).Error
}

