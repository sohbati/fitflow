package repository

import (
	"context"
	"fitflow-business/internal/model"
)

// TrainerRepository defines the interface for trainer database operations
type TrainerRepository interface {
	CreateTrainer(ctx context.Context, trainer *model.Trainer) error
	GetTrainerByID(ctx context.Context, id int64) (*model.Trainer, error)
	GetTrainerByPersonID(ctx context.Context, personID int64) (*model.Trainer, error)
	GetTrainerByUserID(ctx context.Context, userID string) (*model.Trainer, error)
	GetTrainers(ctx context.Context, limit, offset int) ([]*model.Trainer, error)
	UpdateTrainer(ctx context.Context, trainer *model.Trainer) error
	DeleteTrainer(ctx context.Context, id int64) error
	GetRegisteredTrainers(ctx context.Context, limit, offset int) ([]*model.Trainer, error)
}
