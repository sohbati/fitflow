package model

import (
	"time"
	"github.com/google/uuid"
)

// Gender represents the gender of a person
type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
	GenderOther  Gender = "other"
)

// Person represents a base person entity that all roles inherit from
type Person struct {
	ID                      int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID                  uuid.UUID  `json:"user_id" gorm:"type:uuid;not null"`
	FirstName               string     `json:"first_name" gorm:"type:varchar(100);not null"`
	LastName                string     `json:"last_name" gorm:"type:varchar(100);not null"`
	Email                   *string    `json:"email" gorm:"type:varchar(100);unique"`
	PhoneNumber             *string    `json:"phone_number" gorm:"type:varchar(30)"`
	DateOfBirth             *time.Time `json:"date_of_birth" gorm:"type:date"`
	Gender                  *Gender    `json:"gender" gorm:"type:varchar(10)"`
	ProfileImageURL         *string    `json:"profile_image_url" gorm:"type:varchar(500)"`
	Address                 *string    `json:"address" gorm:"type:varchar(255)"`
	City                    *string    `json:"city" gorm:"type:varchar(100)"`
	Province                *string    `json:"province" gorm:"type:varchar(100)"`
	Country                 *string    `json:"country" gorm:"type:varchar(100)"`
	PostalCode              *string    `json:"postal_code" gorm:"type:varchar(20)"`
	EmergencyContactName    *string    `json:"emergency_contact_name" gorm:"type:varchar(150)"`
	EmergencyContactPhone   *string    `json:"emergency_contact_phone" gorm:"type:varchar(30)"`
	EmergencyContactRelation *string   `json:"emergency_contact_relation" gorm:"type:varchar(50)"`
	IsActive                bool       `json:"is_active" gorm:"default:true"`
	CreatedAt               time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt               time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}

// GetFullName returns the full name of the person
func (p *Person) GetFullName() string {
	return p.FirstName + " " + p.LastName
}

// GetAge calculates the age of the person based on date of birth
func (p *Person) GetAge() (int, error) {
	if p.DateOfBirth == nil {
		return 0, nil
	}
	
	now := time.Now()
	age := now.Year() - p.DateOfBirth.Year()
	
	// Adjust if birthday hasn't occurred this year
	if now.YearDay() < p.DateOfBirth.YearDay() {
		age--
	}
	
	return age, nil
}
