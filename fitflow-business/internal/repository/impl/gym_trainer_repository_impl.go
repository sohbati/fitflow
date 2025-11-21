package impl

import (
	"context"
	"fitflow-business/internal/model"
	"fitflow-business/internal/repository"
	"gorm.io/gorm"
)

// gymTrainerRepository implements GymTrainerRepository interface
type gymTrainerRepository struct {
	db *gorm.DB
}

// NewGymTrainerRepository creates a new gym trainer repository
func NewGymTrainerRepository(db *gorm.DB) repository.GymTrainerRepository {
	return &gymTrainerRepository{db: db}
}

// AddTrainerToGym adds a trainer to a gym
func (r *gymTrainerRepository) AddTrainerToGym(ctx context.Context, gymID, trainerID int64) error {
	gymTrainer := &model.GymTrainer{
		GymID:     gymID,
		TrainerID: trainerID,
	}
	return r.db.WithContext(ctx).Create(gymTrainer).Error
}

// RemoveTrainerFromGym removes a trainer from a gym
func (r *gymTrainerRepository) RemoveTrainerFromGym(ctx context.Context, gymID, trainerID int64) error {
	return r.db.WithContext(ctx).Where("gym_id = ? AND trainer_id = ?", gymID, trainerID).
		Delete(&model.GymTrainer{}).Error
}

// GetGymTrainers retrieves all trainers for a gym
func (r *gymTrainerRepository) GetGymTrainers(ctx context.Context, gymID int64) ([]*model.Trainer, error) {
	var trainers []*model.Trainer
	err := r.db.WithContext(ctx).Joins("JOIN gym_trainers ON trainers.id = gym_trainers.trainer_id").
		Where("gym_trainers.gym_id = ?", gymID).Find(&trainers).Error
	return trainers, err
}

// GetTrainerGyms retrieves all gyms for a trainer
func (r *gymTrainerRepository) GetTrainerGyms(ctx context.Context, trainerID int64) ([]*model.Gym, error) {
	var gyms []*model.Gym
	err := r.db.WithContext(ctx).Joins("JOIN gym_trainers ON gyms.id = gym_trainers.gym_id").
		Where("gym_trainers.trainer_id = ?", trainerID).Find(&gyms).Error
	return gyms, err
}
