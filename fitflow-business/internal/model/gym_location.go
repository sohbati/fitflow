package model

import (
	"time"
)

// GymLocation represents a gym location entity
type GymLocation struct {
	ID           int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	GymID        int64     `json:"gym_id" gorm:"not null"`
	LocationType string    `json:"location_type" gorm:"type:varchar(20);not null"` // home, gym, park, nature
	Address      *string   `json:"address" gorm:"type:varchar(255)"`
	City         *string   `json:"city" gorm:"type:varchar(100)"`
	Province     *string   `json:"province" gorm:"type:varchar(100)"`
	Country      *string   `json:"country" gorm:"type:varchar(100)"`
	PostalCode   *string   `json:"postal_code" gorm:"type:varchar(20)"`
	Latitude     *float64  `json:"latitude" gorm:"type:decimal(10,7)"`
	Longitude    *float64  `json:"longitude" gorm:"type:decimal(10,7)"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	
	// Relationships
	Gym Gym `json:"gym,omitempty" gorm:"foreignKey:GymID"`
}
