package impl

import (
	"context"
	"fitflow-business/internal/model"
	"fitflow-business/internal/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// personRepository implements PersonRepository interface
type personRepository struct {
	db *gorm.DB
}

// NewPersonRepository creates a new person repository
func NewPersonRepository(db *gorm.DB) repository.PersonRepository {
	return &personRepository{db: db}
}

// CreatePerson creates a new person
func (r *personRepository) CreatePerson(ctx context.Context, person *model.Person) error {
	return r.db.WithContext(ctx).Create(person).Error
}

// GetPersonByID retrieves a person by ID
func (r *personRepository) GetPersonByID(ctx context.Context, id int64) (*model.Person, error) {
	var person model.Person
	err := r.db.WithContext(ctx).First(&person, id).Error
	if err != nil {
		return nil, err
	}
	return &person, nil
}

// GetPersonByUserID retrieves a person by user ID
func (r *personRepository) GetPersonByUserID(ctx context.Context, userID uuid.UUID) (*model.Person, error) {
	var person model.Person
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&person).Error
	if err != nil {
		return nil, err
	}
	return &person, nil
}

// GetPersons retrieves all persons with pagination
func (r *personRepository) GetPersons(ctx context.Context, limit, offset int) ([]*model.Person, error) {
	var persons []*model.Person
	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&persons).Error
	return persons, err
}

// UpdatePerson updates an existing person
func (r *personRepository) UpdatePerson(ctx context.Context, person *model.Person) error {
	return r.db.WithContext(ctx).Save(person).Error
}

// DeletePerson deletes a person by ID
func (r *personRepository) DeletePerson(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.Person{}, id).Error
}

// SearchPersons searches persons by name, email, or phone
func (r *personRepository) SearchPersons(ctx context.Context, query string, limit, offset int) ([]*model.Person, error) {
	var persons []*model.Person
	err := r.db.WithContext(ctx).
		Where("first_name ILIKE ? OR last_name ILIKE ? OR email ILIKE ? OR phone_number ILIKE ?", 
			"%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%").
		Limit(limit).Offset(offset).Find(&persons).Error
	return persons, err
}

// GetPersonsByEmail retrieves a person by email
func (r *personRepository) GetPersonsByEmail(ctx context.Context, email string) (*model.Person, error) {
	var person model.Person
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&person).Error
	if err != nil {
		return nil, err
	}
	return &person, nil
}

// GetPersonsByPhone retrieves a person by phone number
func (r *personRepository) GetPersonsByPhone(ctx context.Context, phone string) (*model.Person, error) {
	var person model.Person
	err := r.db.WithContext(ctx).Where("phone_number = ?", phone).First(&person).Error
	if err != nil {
		return nil, err
	}
	return &person, nil
}

// GetActivePersons retrieves only active persons
func (r *personRepository) GetActivePersons(ctx context.Context, limit, offset int) ([]*model.Person, error) {
	var persons []*model.Person
	err := r.db.WithContext(ctx).
		Where("is_active = ?", true).
		Limit(limit).Offset(offset).Find(&persons).Error
	return persons, err
}

// GetPersonsByGender retrieves persons by gender
func (r *personRepository) GetPersonsByGender(ctx context.Context, gender model.Gender, limit, offset int) ([]*model.Person, error) {
	var persons []*model.Person
	err := r.db.WithContext(ctx).
		Where("gender = ?", gender).
		Limit(limit).Offset(offset).Find(&persons).Error
	return persons, err
}

// GetPersonsByLocation retrieves persons by location
func (r *personRepository) GetPersonsByLocation(ctx context.Context, city, province, country string, limit, offset int) ([]*model.Person, error) {
	var persons []*model.Person
	query := r.db.WithContext(ctx)
	
	if city != "" {
		query = query.Where("city ILIKE ?", "%"+city+"%")
	}
	if province != "" {
		query = query.Where("province ILIKE ?", "%"+province+"%")
	}
	if country != "" {
		query = query.Where("country ILIKE ?", "%"+country+"%")
	}
	
	err := query.Limit(limit).Offset(offset).Find(&persons).Error
	return persons, err
}

// GetPersonsByUserIDs retrieves persons by multiple user IDs
func (r *personRepository) GetPersonsByUserIDs(ctx context.Context, userIDs []uuid.UUID) ([]*model.Person, error) {
	var persons []*model.Person
	err := r.db.WithContext(ctx).Where("user_id IN ?", userIDs).Find(&persons).Error
	return persons, err
}

// UpdatePersonStatus updates the active status of a person
func (r *personRepository) UpdatePersonStatus(ctx context.Context, id int64, isActive bool) error {
	return r.db.WithContext(ctx).
		Model(&model.Person{}).
		Where("id = ?", id).
		Update("is_active", isActive).Error
}
