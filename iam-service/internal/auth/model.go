package auth

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AuthProvider represents different authentication providers
type AuthProvider struct {
	ID          uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	Name        string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"`
	DisplayName string    `gorm:"type:varchar(100);not null" json:"display_name"`
	Description string    `gorm:"type:text" json:"description"`
	IsActive    bool      `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (AuthProvider) TableName() string {
	return "auth_providers"
}

// BeforeCreate hook to generate UUID
func (ap *AuthProvider) BeforeCreate(tx *gorm.DB) error {
	if ap.ID == uuid.Nil {
		ap.ID = uuid.New()
	}
	return nil
}

// UserAuth links users to their authentication providers
type UserAuth struct {
	ID             uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	UserID         uuid.UUID `gorm:"type:char(36);not null" json:"user_id"`
	AuthProviderID uuid.UUID `gorm:"type:char(36);not null" json:"auth_provider_id"`

	// Provider-specific identifiers
	ProviderUserID string `gorm:"type:varchar(255)" json:"provider_user_id"`
	ProviderEmail  string `gorm:"type:varchar(255)" json:"provider_email"`
	Username       string `gorm:"type:varchar(255)" json:"username"` // For local authentication

	// Provider-specific data (JSON)
	ProviderData string `gorm:"type:json" json:"provider_data"`

	// Verification status
	IsVerified bool       `gorm:"default:false" json:"is_verified"`
	VerifiedAt *time.Time `json:"verified_at"`

	// OTP fields (for mobile/email OTP)
	OTPCode      string     `gorm:"type:varchar(10)" json:"otp_code"`
	OTPExpiresAt *time.Time `json:"otp_expires_at"`
	OTPAttempts  int        `gorm:"default:0" json:"otp_attempts"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relationships (commented to avoid circular imports)
	// User           user.User         `gorm:"foreignKey:UserID" json:"user,omitempty"`
	// AuthProvider   AuthProvider `gorm:"foreignKey:AuthProviderID" json:"auth_provider,omitempty"`
}

func (UserAuth) TableName() string {
	return "user_auth"
}

// BeforeCreate hook to generate UUID
func (ua *UserAuth) BeforeCreate(tx *gorm.DB) error {
	if ua.ID == uuid.Nil {
		ua.ID = uuid.New()
	}
	return nil
}

// Constants for auth provider names
const (
	AuthProviderLocal     = "local"
	AuthProviderGoogle    = "google"
	AuthProviderFacebook  = "facebook"
	AuthProviderApple     = "apple"
	AuthProviderMobileOTP = "mobile_otp"
	AuthProviderEmailOTP  = "email_otp"
)
