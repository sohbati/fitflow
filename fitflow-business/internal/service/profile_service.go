package service

import (
	"context"
	"fitflow-business/internal/model"
	"github.com/google/uuid"
)

// ProfileService defines the interface for profile business logic operations
type ProfileService interface {
	// Basic CRUD operations
	CreateProfile(ctx context.Context, profile *model.Profile) error
	GetProfileByID(ctx context.Context, id int64) (*model.Profile, error)
	GetProfilesByUserID(ctx context.Context, userID uuid.UUID) ([]*model.Profile, error)
	GetProfileByUserIDAndType(ctx context.Context, userID uuid.UUID, profileType model.ProfileType) (*model.Profile, error)
	GetDefaultProfile(ctx context.Context, userID uuid.UUID) (*model.Profile, error)
	UpdateProfile(ctx context.Context, profile *model.Profile) error
	DeleteProfile(ctx context.Context, id int64) error

	// Profile management operations
	SetDefaultProfile(ctx context.Context, userID uuid.UUID, profileID int64) error
	GetActiveProfiles(ctx context.Context, userID uuid.UUID) ([]*model.Profile, error)
	GetProfilesByType(ctx context.Context, profileType model.ProfileType, limit, offset int) ([]*model.Profile, error)
	
	// Profile status operations
	ActivateProfile(ctx context.Context, id int64) error
	DeactivateProfile(ctx context.Context, id int64) error

	// Business logic operations
	CreateGymOwnerProfile(ctx context.Context, userID uuid.UUID, personID int64, gymOwnerID int64) (*model.Profile, error)
	CreateTrainerProfile(ctx context.Context, userID uuid.UUID, personID int64, trainerID int64) (*model.Profile, error)
	CreateTraineeProfile(ctx context.Context, userID uuid.UUID, personID int64, traineeID int64) (*model.Profile, error)
	ValidateProfile(ctx context.Context, profile *model.Profile) error
	
	// Sync profiles from existing role records (for backward compatibility)
	SyncProfilesFromExistingRoles(ctx context.Context, userID uuid.UUID) ([]*model.Profile, error)
}

