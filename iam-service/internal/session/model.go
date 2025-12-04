package session

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	ID             uuid.UUID  `gorm:"type:char(36);primary_key" json:"id"`
	UserID         uuid.UUID  `gorm:"type:char(36);not null" json:"user_id"`
	TokenHash      string     `gorm:"type:varchar(255);uniqueIndex;not null" json:"-"`
	DeviceInfo     string     `gorm:"type:varchar(500)" json:"device_info"`
	IPAddress      string     `gorm:"type:varchar(45)" json:"ip_address"`
	UserAgent      string     `gorm:"type:text" json:"user_agent"`
	ExpiresAt      time.Time  `gorm:"not null" json:"expires_at"`
	LastActivityAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"last_activity_at"`
	CreatedAt      time.Time  `json:"created_at"`
}

func (Session) TableName() string {
	return "sessions"
}

// BeforeCreate hook to generate UUID
func (s *Session) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return nil
}

