package model

import (
	"time"
)

// Gym represents a gym entity
type Gym struct {
	ID           int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string    `json:"name" gorm:"type:varchar(150);not null"`
	Description  *string   `json:"description" gorm:"type:text"`
	PhoneNumber  *string   `json:"phone_number" gorm:"type:varchar(30)"`
	Email        *string   `json:"email" gorm:"type:varchar(100)"`
	WebsiteURL   *string   `json:"website_url" gorm:"type:varchar(255)"`
	IsVerified   bool      `json:"is_verified" gorm:"default:false"`
	Facilities   JSONB     `json:"facilities" gorm:"type:jsonb"`
	Images       Images    `json:"images" gorm:"type:jsonb"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	
	// Relationships
	Locations []GymLocation `json:"locations,omitempty" gorm:"foreignKey:GymID"`
	Owners    []GymOwner    `json:"owners,omitempty" gorm:"foreignKey:GymID"`
	Trainers  []Trainer     `json:"trainers,omitempty" gorm:"many2many:gym_trainers;"`
}