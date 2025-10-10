package auth

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AuthProviderRepository interface for auth provider operations
type AuthProviderRepository interface {
	Create(provider *AuthProvider) error
	GetByID(id uuid.UUID) (*AuthProvider, error)
	GetByName(name string) (*AuthProvider, error)
	GetAll() ([]*AuthProvider, error)
	Update(provider *AuthProvider) error
	Delete(id uuid.UUID) error
}

// UserAuthRepository interface for user auth operations
type UserAuthRepository interface {
	Create(userAuth *UserAuth) error
	GetByID(id uuid.UUID) (*UserAuth, error)
	GetByUserID(userID uuid.UUID) ([]*UserAuth, error)
	GetByProviderAndUserID(providerID uuid.UUID, userID uuid.UUID) (*UserAuth, error)
	GetByProviderAndProviderUserID(providerID uuid.UUID, providerUserID string) (*UserAuth, error)
	GetByProviderAndEmail(providerID uuid.UUID, email string) (*UserAuth, error)
	GetByProviderAndUsername(providerID uuid.UUID, username string) (*UserAuth, error)
	Update(userAuth *UserAuth) error
	Delete(id uuid.UUID) error
	DeleteByUserID(userID uuid.UUID) error
}

// authProviderRepository implements AuthProviderRepository
type authProviderRepository struct {
	db *gorm.DB
}

func NewAuthProviderRepository(db *gorm.DB) AuthProviderRepository {
	return &authProviderRepository{db: db}
}

func (r *authProviderRepository) Create(provider *AuthProvider) error {
	return r.db.Create(provider).Error
}

func (r *authProviderRepository) GetByID(id uuid.UUID) (*AuthProvider, error) {
	var provider AuthProvider
	err := r.db.Where("id = ?", id).First(&provider).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("auth provider not found")
		}
		return nil, err
	}
	return &provider, nil
}

func (r *authProviderRepository) GetByName(name string) (*AuthProvider, error) {
	var provider AuthProvider
	err := r.db.Where("name = ?", name).First(&provider).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("auth provider not found")
		}
		return nil, err
	}
	return &provider, nil
}

func (r *authProviderRepository) GetAll() ([]*AuthProvider, error) {
	var providers []*AuthProvider
	err := r.db.Where("is_active = ?", true).Find(&providers).Error
	return providers, err
}

func (r *authProviderRepository) Update(provider *AuthProvider) error {
	return r.db.Save(provider).Error
}

func (r *authProviderRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&AuthProvider{}, id).Error
}

// userAuthRepository implements UserAuthRepository
type userAuthRepository struct {
	db *gorm.DB
}

func NewUserAuthRepository(db *gorm.DB) UserAuthRepository {
	return &userAuthRepository{db: db}
}

func (r *userAuthRepository) Create(userAuth *UserAuth) error {
	return r.db.Create(userAuth).Error
}

func (r *userAuthRepository) GetByID(id uuid.UUID) (*UserAuth, error) {
	var userAuth UserAuth
	err := r.db.Where("id = ?", id).First(&userAuth).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user auth not found")
		}
		return nil, err
	}
	return &userAuth, nil
}

func (r *userAuthRepository) GetByUserID(userID uuid.UUID) ([]*UserAuth, error) {
	var userAuths []*UserAuth
	err := r.db.Where("user_id = ?", userID).Find(&userAuths).Error
	return userAuths, err
}

func (r *userAuthRepository) GetByProviderAndUserID(providerID uuid.UUID, userID uuid.UUID) (*UserAuth, error) {
	var userAuth UserAuth
	err := r.db.Where("auth_provider_id = ? AND user_id = ?", providerID, userID).First(&userAuth).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user auth not found")
		}
		return nil, err
	}
	return &userAuth, nil
}

func (r *userAuthRepository) GetByProviderAndProviderUserID(providerID uuid.UUID, providerUserID string) (*UserAuth, error) {
	var userAuth UserAuth
	err := r.db.Where("auth_provider_id = ? AND provider_user_id = ?", providerID, providerUserID).First(&userAuth).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user auth not found")
		}
		return nil, err
	}
	return &userAuth, nil
}

func (r *userAuthRepository) GetByProviderAndEmail(providerID uuid.UUID, email string) (*UserAuth, error) {
	var userAuth UserAuth
	err := r.db.Where("auth_provider_id = ? AND provider_email = ?", providerID, email).First(&userAuth).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user auth not found")
		}
		return nil, err
	}
	return &userAuth, nil
}

func (r *userAuthRepository) GetByProviderAndUsername(providerID uuid.UUID, username string) (*UserAuth, error) {
	var userAuth UserAuth
	err := r.db.Where("auth_provider_id = ? AND username = ?", providerID, username).First(&userAuth).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user auth not found")
		}
		return nil, err
	}
	return &userAuth, nil
}

func (r *userAuthRepository) Update(userAuth *UserAuth) error {
	return r.db.Save(userAuth).Error
}

func (r *userAuthRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&UserAuth{}, id).Error
}

func (r *userAuthRepository) DeleteByUserID(userID uuid.UUID) error {
	return r.db.Where("user_id = ?", userID).Delete(&UserAuth{}).Error
}
