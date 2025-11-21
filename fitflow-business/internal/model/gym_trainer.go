package model

import (
	"time"
)

// GymTrainer represents the junction table between gyms and trainers
type GymTrainer struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	GymID     int64     `json:"gym_id" gorm:"not null"`
	TrainerID int64     `json:"trainer_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	// Relationships
	Gym     Gym     `json:"gym,omitempty" gorm:"foreignKey:GymID"`
	Trainer Trainer `json:"trainer,omitempty" gorm:"foreignKey:TrainerID"`
}
