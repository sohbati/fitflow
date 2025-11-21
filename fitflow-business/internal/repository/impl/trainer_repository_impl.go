package impl

import (
	"context"
	"fitflow-business/internal/model"
	"fitflow-business/internal/repository"
	"gorm.io/gorm"
)

// trainerRepository implements TrainerRepository interface
type trainerRepository struct {
	db *gorm.DB
}

// NewTrainerRepository creates a new trainer repository
func NewTrainerRepository(db *gorm.DB) repository.TrainerRepository {
	return &trainerRepository{db: db}
}

// CreateTrainer creates a new trainer
func (r *trainerRepository) CreateTrainer(ctx context.Context, trainer *model.Trainer) error {
	return r.db.WithContext(ctx).Create(trainer).Error
}

// GetTrainerByID retrieves a trainer by ID
func (r *trainerRepository) GetTrainerByID(ctx context.Context, id int64) (*model.Trainer, error) {
	var trainer model.Trainer
	err := r.db.WithContext(ctx).Preload("Gyms").First(&trainer, id).Error
	if err != nil {
		return nil, err
	}
	return &trainer, nil
}

// GetTrainers retrieves all trainers with pagination
func (r *trainerRepository) GetTrainers(ctx context.Context, limit, offset int) ([]*model.Trainer, error) {
	var trainers []*model.Trainer
	err := r.db.WithContext(ctx).Preload("Gyms").Limit(limit).Offset(offset).Find(&trainers).Error
	return trainers, err
}

// UpdateTrainer updates an existing trainer
func (r *trainerRepository) UpdateTrainer(ctx context.Context, trainer *model.Trainer) error {
	return r.db.WithContext(ctx).Save(trainer).Error
}

// DeleteTrainer deletes a trainer by ID
func (r *trainerRepository) DeleteTrainer(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.Trainer{}, id).Error
}

// GetRegisteredTrainers retrieves only registered trainers
func (r *trainerRepository) GetRegisteredTrainers(ctx context.Context, limit, offset int) ([]*model.Trainer, error) {
	var trainers []*model.Trainer
	err := r.db.WithContext(ctx).Preload("Gyms").
		Where("is_registered = ?", true).
		Limit(limit).Offset(offset).Find(&trainers).Error
	return trainers, err
}
