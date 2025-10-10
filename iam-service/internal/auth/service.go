package auth

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// AuthProviderService interface for auth provider business logic
type AuthProviderService interface {
	CreateProvider(provider *AuthProvider) (*AuthProvider, error)
	GetProviderByID(id uuid.UUID) (*AuthProvider, error)
	GetProviderByName(name string) (*AuthProvider, error)
	GetAllProviders() ([]*AuthProvider, error)
	UpdateProvider(provider *AuthProvider) (*AuthProvider, error)
	DeleteProvider(id uuid.UUID) error
	InitializeDefaultProviders() error
}

// UserAuthService interface for user auth business logic
type UserAuthService interface {
	CreateUserAuth(userAuth *UserAuth) (*UserAuth, error)
	GetUserAuthByID(id uuid.UUID) (*UserAuth, error)
	GetUserAuthsByUserID(userID uuid.UUID) ([]*UserAuth, error)
	GetUserAuthByProviderAndUserID(providerID uuid.UUID, userID uuid.UUID) (*UserAuth, error)
	GetUserAuthByProviderAndProviderUserID(providerID uuid.UUID, providerUserID string) (*UserAuth, error)
	GetUserAuthByProviderAndEmail(providerID uuid.UUID, email string) (*UserAuth, error)
	GetUserAuthByProviderAndUsername(providerID uuid.UUID, username string) (*UserAuth, error)
	UpdateUserAuth(userAuth *UserAuth) (*UserAuth, error)
	DeleteUserAuth(id uuid.UUID) error
	DeleteUserAuthsByUserID(userID uuid.UUID) error

	// OTP methods
	GenerateOTP(userID uuid.UUID, providerID uuid.UUID, email string) (*UserAuth, error)
	VerifyOTP(userID uuid.UUID, providerID uuid.UUID, otpCode string) (*UserAuth, error)
}

// authProviderService implements AuthProviderService
type authProviderService struct {
	repo AuthProviderRepository
}

func NewAuthProviderService(repo AuthProviderRepository) AuthProviderService {
	return &authProviderService{repo: repo}
}

func (s *authProviderService) CreateProvider(provider *AuthProvider) (*AuthProvider, error) {
	if err := s.repo.Create(provider); err != nil {
		return nil, err
	}
	return provider, nil
}

func (s *authProviderService) GetProviderByID(id uuid.UUID) (*AuthProvider, error) {
	return s.repo.GetByID(id)
}

func (s *authProviderService) GetProviderByName(name string) (*AuthProvider, error) {
	return s.repo.GetByName(name)
}

func (s *authProviderService) GetAllProviders() ([]*AuthProvider, error) {
	return s.repo.GetAll()
}

func (s *authProviderService) UpdateProvider(provider *AuthProvider) (*AuthProvider, error) {
	if err := s.repo.Update(provider); err != nil {
		return nil, err
	}
	return provider, nil
}

func (s *authProviderService) DeleteProvider(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *authProviderService) InitializeDefaultProviders() error {
	// Check if providers already exist
	providers, err := s.repo.GetAll()
	if err != nil {
		return err
	}

	if len(providers) > 0 {
		return nil // Already initialized
	}

	// Create default providers
	defaultProviders := []*AuthProvider{
		{
			Name:        AuthProviderLocal,
			DisplayName: "Local Authentication",
			Description: "Username/password authentication",
			IsActive:    true,
		},
		{
			Name:        AuthProviderGoogle,
			DisplayName: "Google OAuth",
			Description: "Google OAuth 2.0 authentication",
			IsActive:    true,
		},
		{
			Name:        AuthProviderFacebook,
			DisplayName: "Facebook OAuth",
			Description: "Facebook OAuth 2.0 authentication",
			IsActive:    true,
		},
		{
			Name:        AuthProviderApple,
			DisplayName: "Apple Sign In",
			Description: "Apple Sign In authentication",
			IsActive:    true,
		},
		{
			Name:        AuthProviderMobileOTP,
			DisplayName: "Mobile OTP",
			Description: "Mobile phone OTP authentication",
			IsActive:    true,
		},
		{
			Name:        AuthProviderEmailOTP,
			DisplayName: "Email OTP",
			Description: "Email OTP authentication",
			IsActive:    true,
		},
	}

	for _, provider := range defaultProviders {
		if err := s.repo.Create(provider); err != nil {
			return err
		}
	}

	return nil
}

// userAuthService implements UserAuthService
type userAuthService struct {
	repo         AuthProviderRepository
	userAuthRepo UserAuthRepository
}

func NewUserAuthService(repo AuthProviderRepository, userAuthRepo UserAuthRepository) UserAuthService {
	return &userAuthService{
		repo:         repo,
		userAuthRepo: userAuthRepo,
	}
}

func (s *userAuthService) CreateUserAuth(userAuth *UserAuth) (*UserAuth, error) {
	if err := s.userAuthRepo.Create(userAuth); err != nil {
		return nil, err
	}
	return userAuth, nil
}

func (s *userAuthService) GetUserAuthByID(id uuid.UUID) (*UserAuth, error) {
	return s.userAuthRepo.GetByID(id)
}

func (s *userAuthService) GetUserAuthsByUserID(userID uuid.UUID) ([]*UserAuth, error) {
	return s.userAuthRepo.GetByUserID(userID)
}

func (s *userAuthService) GetUserAuthByProviderAndUserID(providerID uuid.UUID, userID uuid.UUID) (*UserAuth, error) {
	return s.userAuthRepo.GetByProviderAndUserID(providerID, userID)
}

func (s *userAuthService) GetUserAuthByProviderAndProviderUserID(providerID uuid.UUID, providerUserID string) (*UserAuth, error) {
	return s.userAuthRepo.GetByProviderAndProviderUserID(providerID, providerUserID)
}

func (s *userAuthService) GetUserAuthByProviderAndEmail(providerID uuid.UUID, email string) (*UserAuth, error) {
	return s.userAuthRepo.GetByProviderAndEmail(providerID, email)
}

func (s *userAuthService) GetUserAuthByProviderAndUsername(providerID uuid.UUID, username string) (*UserAuth, error) {
	return s.userAuthRepo.GetByProviderAndUsername(providerID, username)
}

func (s *userAuthService) UpdateUserAuth(userAuth *UserAuth) (*UserAuth, error) {
	if err := s.userAuthRepo.Update(userAuth); err != nil {
		return nil, err
	}
	return userAuth, nil
}

func (s *userAuthService) DeleteUserAuth(id uuid.UUID) error {
	return s.userAuthRepo.Delete(id)
}

func (s *userAuthService) DeleteUserAuthsByUserID(userID uuid.UUID) error {
	return s.userAuthRepo.DeleteByUserID(userID)
}

// GenerateOTP generates an OTP for the user
func (s *userAuthService) GenerateOTP(userID uuid.UUID, providerID uuid.UUID, email string) (*UserAuth, error) {
	// Generate 6-digit OTP
	otpCode := generateRandomOTP(6)
	expiresAt := time.Now().Add(10 * time.Minute) // OTP expires in 10 minutes

	// Check if user auth already exists
	existingUserAuth, err := s.userAuthRepo.GetByProviderAndUserID(providerID, userID)
	if err != nil {
		// Create new user auth
		userAuth := &UserAuth{
			UserID:         userID,
			AuthProviderID: providerID,
			ProviderEmail:  email,
			OTPCode:        otpCode,
			OTPExpiresAt:   &expiresAt,
			OTPAttempts:    0,
			IsVerified:     false,
		}
		return s.CreateUserAuth(userAuth)
	}

	// Update existing user auth
	existingUserAuth.OTPCode = otpCode
	existingUserAuth.OTPExpiresAt = &expiresAt
	existingUserAuth.OTPAttempts = 0
	existingUserAuth.IsVerified = false

	return s.UpdateUserAuth(existingUserAuth)
}

// VerifyOTP verifies the OTP code
func (s *userAuthService) VerifyOTP(userID uuid.UUID, providerID uuid.UUID, otpCode string) (*UserAuth, error) {
	userAuth, err := s.userAuthRepo.GetByProviderAndUserID(providerID, userID)
	if err != nil {
		return nil, errors.New("invalid OTP")
	}

	// Check if OTP is expired
	if userAuth.OTPExpiresAt != nil && time.Now().After(*userAuth.OTPExpiresAt) {
		return nil, errors.New("OTP expired")
	}

	// Check if too many attempts
	if userAuth.OTPAttempts >= 3 {
		return nil, errors.New("too many OTP attempts")
	}

	// Increment attempts
	userAuth.OTPAttempts++

	// Verify OTP
	if userAuth.OTPCode != otpCode {
		s.UpdateUserAuth(userAuth) // Update attempts count
		return nil, errors.New("invalid OTP")
	}

	// Mark as verified
	now := time.Now()
	userAuth.IsVerified = true
	userAuth.VerifiedAt = &now
	userAuth.OTPCode = "" // Clear OTP after successful verification

	return s.UpdateUserAuth(userAuth)
}

// Helper function to generate random OTP
func generateRandomOTP(length int) string {
	// Simple OTP generation - in production, use crypto/rand
	chars := "0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[time.Now().UnixNano()%int64(len(chars))]
	}
	return string(result)
}
