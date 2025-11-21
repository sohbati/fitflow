package model

import (
	"time"
)

// GymOwner represents a gym owner entity that inherits from Person
type GymOwner struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	PersonID  int64     `json:"person_id" gorm:"not null"`
	GymID     int64     `json:"gym_id" gorm:"not null"`
	BriefBio  *string   `json:"brief_bio" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	Person Person `json:"person,omitempty" gorm:"foreignKey:PersonID"`
	Gym    Gym    `json:"gym,omitempty" gorm:"foreignKey:GymID"`
}
