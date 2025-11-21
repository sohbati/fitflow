package model

import (
	"time"
)

// Trainer represents a trainer entity that inherits from Person
type Trainer struct {
	ID           int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	PersonID     int64     `json:"person_id" gorm:"not null"`
	IsRegistered bool      `json:"is_registered" gorm:"default:false"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	Person Person `json:"person,omitempty" gorm:"foreignKey:PersonID"`
	Gyms   []Gym  `json:"gyms,omitempty" gorm:"many2many:gym_trainers;"`
}
