package role

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(role *Role) error
	GetByRole(role string) (*Role, error)
	GetByID(id uuid.UUID) (*Role, error)
	GetAll() ([]*Role, error)
	Update(role *Role) error
	Delete(id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(role *Role) error {
	return r.db.Create(role).Error
}

func (r *repository) GetByRole(role string) (*Role, error) {
	var roleModel Role
	err := r.db.Where("role = ?", role).First(&roleModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("role not found")
		}
		return nil, err
	}
	return &roleModel, nil
}

func (r *repository) GetByID(id uuid.UUID) (*Role, error) {
	var role Role
	err := r.db.Where("id = ?", id).First(&role).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("role not found")
		}
		return nil, err
	}
	return &role, nil
}

func (r *repository) GetAll() ([]*Role, error) {
	var roles []*Role
	err := r.db.Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *repository) Update(role *Role) error {
	return r.db.Save(role).Error
}

func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&Role{}, id).Error
}
