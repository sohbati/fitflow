package model

import (
	"time"
)

// FitnessLevel represents the fitness level of a trainee
type FitnessLevel string

const (
	FitnessLevelBeginner     FitnessLevel = "beginner"
	FitnessLevelIntermediate FitnessLevel = "intermediate"
	FitnessLevelAdvanced     FitnessLevel = "advanced"
)

// MembershipType represents the type of membership
type MembershipType string

const (
	MembershipTypeBasic   MembershipType = "basic"
	MembershipTypePremium MembershipType = "premium"
	MembershipTypeVIP     MembershipType = "vip"
)

// Trainee represents a trainee entity that inherits from Person
type Trainee struct {
	ID                  int64          `json:"id" gorm:"primaryKey;autoIncrement"`
	PersonID            int64          `json:"person_id" gorm:"not null"`
	HeightCm            *int           `json:"height_cm" gorm:"type:integer"`
	WeightKg            *float64       `json:"weight_kg" gorm:"type:decimal(5,2)"`
	FitnessLevel        FitnessLevel   `json:"fitness_level" gorm:"type:varchar(20);default:'beginner'"`
	Goals               *string        `json:"goals" gorm:"type:text"`
	MedicalConditions   *string        `json:"medical_conditions" gorm:"type:text"`
	MembershipType      MembershipType `json:"membership_type" gorm:"type:varchar(50);default:'basic'"`
	MembershipStartDate *time.Time     `json:"membership_start_date" gorm:"type:date"`
	MembershipEndDate   *time.Time     `json:"membership_end_date" gorm:"type:date"`
	IsActive            bool           `json:"is_active" gorm:"default:true"`
	CreatedAt           time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt           time.Time      `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	Person Person `json:"person,omitempty" gorm:"foreignKey:PersonID"`
}
