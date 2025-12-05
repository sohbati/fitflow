package model

import (
	"github.com/google/uuid"
	"time"
)

// ProfileType represents the type of profile
type ProfileType string

const (
	ProfileTypeGymOwner ProfileType = "gym_owner"
	ProfileTypeTrainer  ProfileType = "trainer"
	ProfileTypeTrainee  ProfileType = "trainee"
)

// Profile represents a user's profile in the system
// A user can have multiple profiles (gym owner, trainer, trainee)
type Profile struct {
	ID        int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uuid.UUID  `json:"user_id" gorm:"type:uuid;not null;index"`
	Type      ProfileType `json:"type" gorm:"type:varchar(20);not null;index"`
	PersonID  int64      `json:"person_id" gorm:"not null;index"`
	
	// Optional references to role-specific entities
	// These will be populated based on the profile type
	GymOwnerID *int64 `json:"gym_owner_id,omitempty" gorm:"index"`
	TrainerID  *int64 `json:"trainer_id,omitempty" gorm:"index"`
	TraineeID  *int64 `json:"trainee_id,omitempty" gorm:"index"`
	
	IsActive  bool      `json:"is_active" gorm:"default:true;index"`
	IsDefault bool      `json:"is_default" gorm:"default:false;index"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	Person Person `json:"person,omitempty" gorm:"foreignKey:PersonID"`
}

// TableName specifies the table name for GORM
func (Profile) TableName() string {
	return "profiles"
}

// IsValidProfileType checks if the profile type is valid
func IsValidProfileType(profileType ProfileType) bool {
	return profileType == ProfileTypeGymOwner ||
		profileType == ProfileTypeTrainer ||
		profileType == ProfileTypeTrainee
}

