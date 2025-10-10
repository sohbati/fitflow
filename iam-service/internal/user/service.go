package user

import (
	"errors"

	"github.com/google/uuid"
)

type Service interface {
	CreateUser(req CreateUserRequest) (*User, error)
	CreateUserFromStruct(user *User) (*User, error)
	AuthenticateUser(username, password string) (*User, error)
	GetUserByID(id uuid.UUID) (*User, error)
	GetUserByEmail(email string) (*User, error)
	UpdateUser(user *User) (*User, error)
}

type service struct {
	repo Repository
}

type CreateUserRequest struct {
	Email       string `json:"email" binding:"required,email"`
	DisplayName string `json:"display_name" binding:"required"`
	Country     string `json:"country" binding:"required"`
	Role        string `json:"role,omitempty"`
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateUser(req CreateUserRequest) (*User, error) {
	// Check if email already exists
	existingUserByEmail, _ := s.repo.GetByEmail(req.Email)
	if existingUserByEmail != nil {
		return nil, errors.New("email already exists")
	}

	// Set default role if not provided
	role := req.Role
	if role == "" {
		role = "user"
	}

	user := &User{
		Email:       req.Email,
		DisplayName: req.DisplayName,
		Country:     req.Country,
		Role:        role,
		IsActive:    true,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// AuthenticateUser is deprecated - use LocalAuthHandler instead
func (s *service) AuthenticateUser(username, password string) (*User, error) {
	return nil, errors.New("use LocalAuthHandler for username/password authentication")
}

func (s *service) GetUserByID(id uuid.UUID) (*User, error) {
	return s.repo.GetByID(id)
}

// CreateUserFromStruct creates a user from User struct (for OAuth)
func (s *service) CreateUserFromStruct(user *User) (*User, error) {
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByEmail gets user by email
func (s *service) GetUserByEmail(email string) (*User, error) {
	return s.repo.GetByEmail(email)
}

// UpdateUser updates an existing user
func (s *service) UpdateUser(user *User) (*User, error) {
	if err := s.repo.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}
