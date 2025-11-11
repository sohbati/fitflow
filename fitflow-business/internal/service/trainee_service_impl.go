package service

import (
	"context"
	"errors"
	"fitflow-business/internal/model"
	"fitflow-business/internal/repository"
	"time"
)

// traineeService implements TraineeService interface
type traineeService struct {
	traineeRepo repository.TraineeRepository
}

// NewTraineeService creates a new trainee service
func NewTraineeService(traineeRepo repository.TraineeRepository) TraineeService {
	return &traineeService{
		traineeRepo: traineeRepo,
	}
}

// CreateTrainee creates a new trainee with validation
func (s *traineeService) CreateTrainee(ctx context.Context, trainee *model.Trainee) error {
	// Validate required fields
	if trainee.Name == "" {
		return errors.New("trainee name is required")
	}
	
	// Validate email uniqueness if provided
	if trainee.Email != nil && *trainee.Email != "" {
		existingTrainee, err := s.traineeRepo.GetTraineesByEmail(ctx, *trainee.Email)
		if err == nil && existingTrainee != nil {
			return errors.New("email already exists")
		}
	}
	
	// Validate phone uniqueness if provided
	if trainee.PhoneNumber != nil && *trainee.PhoneNumber != "" {
		existingTrainee, err := s.traineeRepo.GetTraineesByPhone(ctx, *trainee.PhoneNumber)
		if err == nil && existingTrainee != nil {
			return errors.New("phone number already exists")
		}
	}
	
	// Set default values
	if trainee.FitnessLevel == "" {
		trainee.FitnessLevel = model.FitnessLevelBeginner
	}
	if trainee.MembershipType == "" {
		trainee.MembershipType = model.MembershipTypeBasic
	}
	
	return s.traineeRepo.CreateTrainee(ctx, trainee)
}

// GetTraineeByID retrieves a trainee by ID
func (s *traineeService) GetTraineeByID(ctx context.Context, id int64) (*model.Trainee, error) {
	if id <= 0 {
		return nil, errors.New("invalid trainee ID")
	}
	return s.traineeRepo.GetTraineeByID(ctx, id)
}

// GetTrainees retrieves all trainees with pagination
func (s *traineeService) GetTrainees(ctx context.Context, limit, offset int) ([]*model.Trainee, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}
	
	return s.traineeRepo.GetTrainees(ctx, limit, offset)
}

// UpdateTrainee updates an existing trainee
func (s *traineeService) UpdateTrainee(ctx context.Context, trainee *model.Trainee) error {
	if trainee.ID <= 0 {
		return errors.New("invalid trainee ID")
	}
	if trainee.Name == "" {
		return errors.New("trainee name is required")
	}
	
	// Validate email uniqueness if provided
	if trainee.Email != nil && *trainee.Email != "" {
		existingTrainee, err := s.traineeRepo.GetTraineesByEmail(ctx, *trainee.Email)
		if err == nil && existingTrainee != nil && existingTrainee.ID != trainee.ID {
			return errors.New("email already exists")
		}
	}
	
	// Validate phone uniqueness if provided
	if trainee.PhoneNumber != nil && *trainee.PhoneNumber != "" {
		existingTrainee, err := s.traineeRepo.GetTraineesByPhone(ctx, *trainee.PhoneNumber)
		if err == nil && existingTrainee != nil && existingTrainee.ID != trainee.ID {
			return errors.New("phone number already exists")
		}
	}
	
	return s.traineeRepo.UpdateTrainee(ctx, trainee)
}

// DeleteTrainee deletes a trainee by ID
func (s *traineeService) DeleteTrainee(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid trainee ID")
	}
	return s.traineeRepo.DeleteTrainee(ctx, id)
}

// SearchTrainees searches trainees by name, email, or phone
func (s *traineeService) SearchTrainees(ctx context.Context, query string, limit, offset int) ([]*model.Trainee, error) {
	if query == "" {
		return nil, errors.New("search query cannot be empty")
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
	
	return s.traineeRepo.SearchTrainees(ctx, query, limit, offset)
}

// GetTraineesByEmail retrieves a trainee by email
func (s *traineeService) GetTraineesByEmail(ctx context.Context, email string) (*model.Trainee, error) {
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}
	return s.traineeRepo.GetTraineesByEmail(ctx, email)
}

// GetTraineesByPhone retrieves a trainee by phone number
func (s *traineeService) GetTraineesByPhone(ctx context.Context, phone string) (*model.Trainee, error) {
	if phone == "" {
		return nil, errors.New("phone number cannot be empty")
	}
	return s.traineeRepo.GetTraineesByPhone(ctx, phone)
}

// GetActiveTrainees retrieves only active trainees
func (s *traineeService) GetActiveTrainees(ctx context.Context, limit, offset int) ([]*model.Trainee, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}
	
	return s.traineeRepo.GetActiveTrainees(ctx, limit, offset)
}

// GetTraineesByMembershipType retrieves trainees by membership type
func (s *traineeService) GetTraineesByMembershipType(ctx context.Context, membershipType model.MembershipType, limit, offset int) ([]*model.Trainee, error) {
	if membershipType == "" {
		return nil, errors.New("membership type cannot be empty")
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
	
	return s.traineeRepo.GetTraineesByMembershipType(ctx, membershipType, limit, offset)
}

// GetTraineesByFitnessLevel retrieves trainees by fitness level
func (s *traineeService) GetTraineesByFitnessLevel(ctx context.Context, fitnessLevel model.FitnessLevel, limit, offset int) ([]*model.Trainee, error) {
	if fitnessLevel == "" {
		return nil, errors.New("fitness level cannot be empty")
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
	
	return s.traineeRepo.GetTraineesByFitnessLevel(ctx, fitnessLevel, limit, offset)
}

// GetExpiringMemberships retrieves trainees with memberships expiring within specified days
func (s *traineeService) GetExpiringMemberships(ctx context.Context, days int, limit, offset int) ([]*model.Trainee, error) {
	if days < 0 {
		return nil, errors.New("days cannot be negative")
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
	
	return s.traineeRepo.GetExpiringMemberships(ctx, days, limit, offset)
}

// GetExpiredMemberships retrieves trainees with expired memberships
func (s *traineeService) GetExpiredMemberships(ctx context.Context, limit, offset int) ([]*model.Trainee, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}
	
	return s.traineeRepo.GetExpiredMemberships(ctx, limit, offset)
}

// UpdateMembershipStatus updates the active status of a trainee
func (s *traineeService) UpdateMembershipStatus(ctx context.Context, id int64, isActive bool) error {
	if id <= 0 {
		return errors.New("invalid trainee ID")
	}
	return s.traineeRepo.UpdateMembershipStatus(ctx, id, isActive)
}

// CalculateAge calculates the age of a trainee based on date of birth
func (s *traineeService) CalculateAge(ctx context.Context, trainee *model.Trainee) (int, error) {
	if trainee.DateOfBirth == nil {
		return 0, errors.New("date of birth is required to calculate age")
	}
	
	now := time.Now()
	age := now.Year() - trainee.DateOfBirth.Year()
	
	// Adjust if birthday hasn't occurred this year
	if now.YearDay() < trainee.DateOfBirth.YearDay() {
		age--
	}
	
	return age, nil
}

// CalculateBMI calculates the BMI of a trainee
func (s *traineeService) CalculateBMI(ctx context.Context, trainee *model.Trainee) (float64, error) {
	if trainee.HeightCm == nil || trainee.WeightKg == nil {
		return 0, errors.New("height and weight are required to calculate BMI")
	}
	
	if *trainee.HeightCm <= 0 || *trainee.WeightKg <= 0 {
		return 0, errors.New("height and weight must be positive values")
	}
	
	// Convert height from cm to meters
	heightM := float64(*trainee.HeightCm) / 100.0
	
	// Calculate BMI: weight(kg) / height(m)^2
	bmi := *trainee.WeightKg / (heightM * heightM)
	
	return bmi, nil
}

// ValidateMembership validates if a trainee's membership is active
func (s *traineeService) ValidateMembership(ctx context.Context, trainee *model.Trainee) (bool, error) {
	if !trainee.IsActive {
		return false, nil
	}
	
	if trainee.MembershipEndDate == nil {
		return true, nil // No end date means lifetime membership
	}
	
	now := time.Now()
	return trainee.MembershipEndDate.After(now), nil
}
