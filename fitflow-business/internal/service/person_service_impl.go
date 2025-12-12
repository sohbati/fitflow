package service

import (
	"context"
	"errors"
	"fitflow-business/internal/model"
	"fitflow-business/internal/repository"

	"github.com/google/uuid"
)

// personService implements PersonService interface
type personService struct {
	personRepo repository.PersonRepository
}

// NewPersonService creates a new person service
func NewPersonService(personRepo repository.PersonRepository) PersonService {
	return &personService{
		personRepo: personRepo,
	}
}

// CreatePerson creates a new person with validation
func (s *personService) CreatePerson(ctx context.Context, person *model.Person) error {
	// Validate required fields
	if person.FirstName == "" {
		return errors.New("first_name_is_required")
	}
	if person.LastName == "" {
		return errors.New("last_name_is_required")
	}
	if person.UserID == uuid.Nil {
		return errors.New("user_id_is_required")
	}

	// Validate email uniqueness if provided
	if person.Email != nil && *person.Email != "" {
		existingPerson, err := s.personRepo.GetPersonsByEmail(ctx, *person.Email)
		if err == nil && existingPerson != nil {
			return errors.New("email_already_exists")
		}
	}

	// Validate phone uniqueness if provided
	if person.PhoneNumber != nil && *person.PhoneNumber != "" {
		existingPerson, err := s.personRepo.GetPersonsByPhone(ctx, *person.PhoneNumber)
		if err == nil && existingPerson != nil {
			return errors.New("phone_number_already_exists")
		}
	}

	// Validate user ID uniqueness
	existingPerson, err := s.personRepo.GetPersonByUserID(ctx, person.UserID)
	if err == nil && existingPerson != nil {
		return errors.New("person_with_this_user_id_already_exists")
	}

	return s.personRepo.CreatePerson(ctx, person)
}

// GetPersonByID retrieves a person by ID
func (s *personService) GetPersonByID(ctx context.Context, id int64) (*model.Person, error) {
	if id <= 0 {
		return nil, errors.New("invalid_person_id")
	}
	return s.personRepo.GetPersonByID(ctx, id)
}

// GetPersonByUserID retrieves a person by user ID
func (s *personService) GetPersonByUserID(ctx context.Context, userID uuid.UUID) (*model.Person, error) {
	if userID == uuid.Nil {
		return nil, errors.New("invalid_user_id")
	}
	return s.personRepo.GetPersonByUserID(ctx, userID)
}

// GetPersons retrieves all persons with pagination
func (s *personService) GetPersons(ctx context.Context, limit, offset int) ([]*model.Person, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	return s.personRepo.GetPersons(ctx, limit, offset)
}

// UpdatePerson updates an existing person
func (s *personService) UpdatePerson(ctx context.Context, person *model.Person) error {
	if person.ID <= 0 {
		return errors.New("invalid_person_id")
	}
	if person.FirstName == "" {
		return errors.New("first_name_is_required")
	}
	if person.LastName == "" {
		return errors.New("last_name_is_required")
	}

	// Validate email uniqueness if provided
	if person.Email != nil && *person.Email != "" {
		existingPerson, err := s.personRepo.GetPersonsByEmail(ctx, *person.Email)
		if err == nil && existingPerson != nil && existingPerson.ID != person.ID {
			return errors.New("email_already_exists")
		}
	}

	// Validate phone uniqueness if provided
	if person.PhoneNumber != nil && *person.PhoneNumber != "" {
		existingPerson, err := s.personRepo.GetPersonsByPhone(ctx, *person.PhoneNumber)
		if err == nil && existingPerson != nil && existingPerson.ID != person.ID {
			return errors.New("phone_number_already_exists")
		}
	}

	return s.personRepo.UpdatePerson(ctx, person)
}

// DeletePerson deletes a person by ID
func (s *personService) DeletePerson(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid_person_id")
	}
	return s.personRepo.DeletePerson(ctx, id)
}

// SearchPersons searches persons by name, email, or phone
func (s *personService) SearchPersons(ctx context.Context, query string, limit, offset int) ([]*model.Person, error) {
	if query == "" {
		return nil, errors.New("search_query_cannot_be_empty")
	}
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	return s.personRepo.SearchPersons(ctx, query, limit, offset)
}

// GetPersonsByEmail retrieves a person by email
func (s *personService) GetPersonsByEmail(ctx context.Context, email string) (*model.Person, error) {
	if email == "" {
		return nil, errors.New("email_cannot_be_empty")
	}
	return s.personRepo.GetPersonsByEmail(ctx, email)
}

// GetPersonsByPhone retrieves a person by phone number
func (s *personService) GetPersonsByPhone(ctx context.Context, phone string) (*model.Person, error) {
	if phone == "" {
		return nil, errors.New("phone_number_cannot_be_empty")
	}
	return s.personRepo.GetPersonsByPhone(ctx, phone)
}

// GetActivePersons retrieves only active persons
func (s *personService) GetActivePersons(ctx context.Context, limit, offset int) ([]*model.Person, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	return s.personRepo.GetActivePersons(ctx, limit, offset)
}

// GetPersonsByGender retrieves persons by gender
func (s *personService) GetPersonsByGender(ctx context.Context, gender model.Gender, limit, offset int) ([]*model.Person, error) {
	if gender == "" {
		return nil, errors.New("gender_cannot_be_empty")
	}
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	return s.personRepo.GetPersonsByGender(ctx, gender, limit, offset)
}

// GetPersonsByLocation retrieves persons by location
func (s *personService) GetPersonsByLocation(ctx context.Context, city, province, country string, limit, offset int) ([]*model.Person, error) {
	if city == "" && province == "" && country == "" {
		return nil, errors.New("at_least_one_location_parameter_is_required")
	}
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	return s.personRepo.GetPersonsByLocation(ctx, city, province, country, limit, offset)
}

// GetPersonsByUserIDs retrieves persons by multiple user IDs
func (s *personService) GetPersonsByUserIDs(ctx context.Context, userIDs []uuid.UUID) ([]*model.Person, error) {
	if len(userIDs) == 0 {
		return nil, errors.New("user_ids_list_cannot_be_empty")
	}
	return s.personRepo.GetPersonsByUserIDs(ctx, userIDs)
}

// UpdatePersonStatus updates the active status of a person
func (s *personService) UpdatePersonStatus(ctx context.Context, id int64, isActive bool) error {
	if id <= 0 {
		return errors.New("invalid_person_id")
	}
	return s.personRepo.UpdatePersonStatus(ctx, id, isActive)
}

// CalculateAge calculates the age of a person based on date of birth
func (s *personService) CalculateAge(ctx context.Context, person *model.Person) (int, error) {
	return person.GetAge()
}

// GetFullName returns the full name of a person
func (s *personService) GetFullName(ctx context.Context, person *model.Person) string {
	return person.GetFullName()
}

// ValidatePersonData validates person data
func (s *personService) ValidatePersonData(ctx context.Context, person *model.Person) error {
	if person.FirstName == "" {
		return errors.New("first_name_is_required")
	}
	if person.LastName == "" {
		return errors.New("last_name_is_required")
	}
	if person.UserID == uuid.Nil {
		return errors.New("user_id_is_required")
	}

	// Validate email format if provided
	if person.Email != nil && *person.Email != "" {
		// Basic email validation
		if len(*person.Email) < 5 || !contains(*person.Email, "@") {
			return errors.New("invalid_email_format")
		}
	}

	// Validate phone format if provided
	if person.PhoneNumber != nil && *person.PhoneNumber != "" {
		// Basic phone validation
		if len(*person.PhoneNumber) < 10 {
			return errors.New("phone_number_must_be_at_least_10_characters")
		}
	}

	return nil
}

// Helper function to check if string contains substring
func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
