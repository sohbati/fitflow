package repository

import (
	"context"
	"fitflow-business/internal/model"
	"github.com/google/uuid"
)

// PersonRepository defines the interface for person database operations
type PersonRepository interface {
	// Basic CRUD operations
	CreatePerson(ctx context.Context, person *model.Person) error
	GetPersonByID(ctx context.Context, id int64) (*model.Person, error)
	GetPersonByUserID(ctx context.Context, userID uuid.UUID) (*model.Person, error)
	GetPersons(ctx context.Context, limit, offset int) ([]*model.Person, error)
	UpdatePerson(ctx context.Context, person *model.Person) error
	DeletePerson(ctx context.Context, id int64) error
	
	// Search and filtering operations
	SearchPersons(ctx context.Context, query string, limit, offset int) ([]*model.Person, error)
	GetPersonsByEmail(ctx context.Context, email string) (*model.Person, error)
	GetPersonsByPhone(ctx context.Context, phone string) (*model.Person, error)
	GetActivePersons(ctx context.Context, limit, offset int) ([]*model.Person, error)
	GetPersonsByGender(ctx context.Context, gender model.Gender, limit, offset int) ([]*model.Person, error)
	GetPersonsByLocation(ctx context.Context, city, province, country string, limit, offset int) ([]*model.Person, error)
	
	// User relationship operations
	GetPersonsByUserIDs(ctx context.Context, userIDs []uuid.UUID) ([]*model.Person, error)
	UpdatePersonStatus(ctx context.Context, id int64, isActive bool) error
}
