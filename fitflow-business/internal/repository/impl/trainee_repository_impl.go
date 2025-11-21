package impl

import (
	"context"
	"fitflow-business/internal/model"
	"fitflow-business/internal/repository"
	"gorm.io/gorm"
)

// traineeRepository implements TraineeRepository interface
type traineeRepository struct {
	db *gorm.DB
}

// NewTraineeRepository creates a new trainee repository
func NewTraineeRepository(db *gorm.DB) repository.TraineeRepository {
	return &traineeRepository{db: db}
}

// CreateTrainee creates a new trainee
func (r *traineeRepository) CreateTrainee(ctx context.Context, trainee *model.Trainee) error {
	return r.db.WithContext(ctx).Create(trainee).Error
}

// GetTraineeByID retrieves a trainee by ID
func (r *traineeRepository) GetTraineeByID(ctx context.Context, id int64) (*model.Trainee, error) {
	var trainee model.Trainee
	err := r.db.WithContext(ctx).Preload("Person").First(&trainee, id).Error
	if err != nil {
		return nil, err
	}
	return &trainee, nil
}

// GetTrainees retrieves all trainees with pagination
func (r *traineeRepository) GetTrainees(ctx context.Context, limit, offset int) ([]*model.Trainee, error) {
	var trainees []*model.Trainee
	err := r.db.WithContext(ctx).
		Preload("Person").
		Limit(limit).
		Offset(offset).
		Find(&trainees).Error
	return trainees, err
}

// UpdateTrainee updates an existing trainee
func (r *traineeRepository) UpdateTrainee(ctx context.Context, trainee *model.Trainee) error {
	return r.db.WithContext(ctx).Save(trainee).Error
}

// DeleteTrainee deletes a trainee by ID
func (r *traineeRepository) DeleteTrainee(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.Trainee{}, id).Error
}

// SearchTrainees searches trainees by name, email, or phone
func (r *traineeRepository) SearchTrainees(ctx context.Context, query string, limit, offset int) ([]*model.Trainee, error) {
	var trainees []*model.Trainee
	like := "%" + query + "%"
	err := r.db.WithContext(ctx).
		Model(&model.Trainee{}).
		Preload("Person").
		Joins("Person").
		Where("persons.first_name ILIKE ? OR persons.last_name ILIKE ? OR persons.email ILIKE ? OR persons.phone_number ILIKE ?",
			like, like, like, like).
		Limit(limit).Offset(offset).Find(&trainees).Error
	return trainees, err
}

// GetTraineesByEmail retrieves a trainee by email
func (r *traineeRepository) GetTraineesByEmail(ctx context.Context, email string) (*model.Trainee, error) {
	var trainee model.Trainee
	err := r.db.WithContext(ctx).
		Preload("Person").
		Joins("Person").
		Where("persons.email = ?", email).
		First(&trainee).Error
	if err != nil {
		return nil, err
	}
	return &trainee, nil
}

// GetTraineesByPhone retrieves a trainee by phone number
func (r *traineeRepository) GetTraineesByPhone(ctx context.Context, phone string) (*model.Trainee, error) {
	var trainee model.Trainee
	err := r.db.WithContext(ctx).
		Preload("Person").
		Joins("Person").
		Where("persons.phone_number = ?", phone).
		First(&trainee).Error
	if err != nil {
		return nil, err
	}
	return &trainee, nil
}

// GetActiveTrainees retrieves only active trainees
func (r *traineeRepository) GetActiveTrainees(ctx context.Context, limit, offset int) ([]*model.Trainee, error) {
	var trainees []*model.Trainee
	err := r.db.WithContext(ctx).
		Preload("Person").
		Where("is_active = ?", true).
		Limit(limit).Offset(offset).Find(&trainees).Error
	return trainees, err
}

// GetTraineesByMembershipType retrieves trainees by membership type
func (r *traineeRepository) GetTraineesByMembershipType(ctx context.Context, membershipType model.MembershipType, limit, offset int) ([]*model.Trainee, error) {
	var trainees []*model.Trainee
	err := r.db.WithContext(ctx).
		Preload("Person").
		Where("membership_type = ?", membershipType).
		Limit(limit).Offset(offset).Find(&trainees).Error
	return trainees, err
}

// GetTraineesByFitnessLevel retrieves trainees by fitness level
func (r *traineeRepository) GetTraineesByFitnessLevel(ctx context.Context, fitnessLevel model.FitnessLevel, limit, offset int) ([]*model.Trainee, error) {
	var trainees []*model.Trainee
	err := r.db.WithContext(ctx).
		Preload("Person").
		Where("fitness_level = ?", fitnessLevel).
		Limit(limit).Offset(offset).Find(&trainees).Error
	return trainees, err
}

// GetExpiringMemberships retrieves trainees with memberships expiring within specified days
func (r *traineeRepository) GetExpiringMemberships(ctx context.Context, days int, limit, offset int) ([]*model.Trainee, error) {
	var trainees []*model.Trainee
	err := r.db.WithContext(ctx).
		Preload("Person").
		Where("membership_end_date BETWEEN ? AND ? AND is_active = ?",
			gorm.Expr("CURRENT_DATE"),
			gorm.Expr("CURRENT_DATE + INTERVAL ? DAY", days),
			true).
		Limit(limit).Offset(offset).Find(&trainees).Error
	return trainees, err
}

// GetExpiredMemberships retrieves trainees with expired memberships
func (r *traineeRepository) GetExpiredMemberships(ctx context.Context, limit, offset int) ([]*model.Trainee, error) {
	var trainees []*model.Trainee
	err := r.db.WithContext(ctx).
		Preload("Person").
		Where("membership_end_date < CURRENT_DATE AND is_active = ?", true).
		Limit(limit).Offset(offset).Find(&trainees).Error
	return trainees, err
}

// UpdateMembershipStatus updates the active status of a trainee
func (r *traineeRepository) UpdateMembershipStatus(ctx context.Context, id int64, isActive bool) error {
	return r.db.WithContext(ctx).
		Model(&model.Trainee{}).
		Where("id = ?", id).
		Update("is_active", isActive).Error
}
