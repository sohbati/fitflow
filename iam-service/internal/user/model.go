package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	Email       string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	DisplayName string    `gorm:"type:varchar(255);not null" json:"display_name"`
	AvatarURL   string    `gorm:"type:varchar(500)" json:"avatar_url"`
	Country     string    `gorm:"type:varchar(255)" json:"country"`
	Role        string    `gorm:"type:varchar(255);default:user" json:"-"` // Hidden from JSON responses
	IsActive    bool      `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relationships (will be loaded separately to avoid circular imports)
	// UserAuths []UserAuth `gorm:"foreignKey:UserID" json:"-"` // Hidden from JSON responses
}

func (User) TableName() string {
	return "users"
}

// BeforeCreate hook to set default role and generate UUID
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	if u.Role == "" {
		u.Role = "user"
	}
	return nil
}
