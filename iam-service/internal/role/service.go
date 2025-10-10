package role

import (
	"errors"

	"github.com/google/uuid"
)

type Service interface {
	CreateRole(req CreateRoleRequest) (*Role, error)
	GetRoleByRole(role string) (*Role, error)
	GetRoleByID(id uuid.UUID) (*Role, error)
	GetAllRoles() ([]*Role, error)
	UpdateRole(id uuid.UUID, req UpdateRoleRequest) (*Role, error)
	DeleteRole(id uuid.UUID) error
}

type service struct {
	repo Repository
}

type CreateRoleRequest struct {
	Role        string `json:"role" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateRoleRequest struct {
	Role        string `json:"role,omitempty"`
	Description string `json:"description,omitempty"`
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateRole(req CreateRoleRequest) (*Role, error) {
	// Check if role already exists
	existingRole, _ := s.repo.GetByRole(req.Role)
	if existingRole != nil {
		return nil, errors.New("role already exists")
	}

	role := &Role{
		Role:        req.Role,
		Description: req.Description,
	}

	if err := s.repo.Create(role); err != nil {
		return nil, err
	}

	return role, nil
}

func (s *service) GetRoleByRole(role string) (*Role, error) {
	return s.repo.GetByRole(role)
}

func (s *service) GetRoleByID(id uuid.UUID) (*Role, error) {
	return s.repo.GetByID(id)
}

func (s *service) GetAllRoles() ([]*Role, error) {
	return s.repo.GetAll()
}

func (s *service) UpdateRole(id uuid.UUID, req UpdateRoleRequest) (*Role, error) {
	role, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Check if new role name conflicts with existing role
	if req.Role != "" && req.Role != role.Role {
		existingRole, _ := s.repo.GetByRole(req.Role)
		if existingRole != nil {
			return nil, errors.New("role already exists")
		}
		role.Role = req.Role
	}

	if req.Description != "" {
		role.Description = req.Description
	}

	if err := s.repo.Update(role); err != nil {
		return nil, err
	}

	return role, nil
}

func (s *service) DeleteRole(id uuid.UUID) error {
	// Check if role exists
	_, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
