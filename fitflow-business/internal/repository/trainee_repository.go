package repository

import (
	"context"
	"fitflow-business/internal/model"
)

// TraineeRepository defines the interface for trainee database operations
type TraineeRepository interface {
	// Basic CRUD operations
	CreateTrainee(ctx context.Context, trainee *model.Trainee) error
	GetTraineeByID(ctx context.Context, id int64) (*model.Trainee, error)
	GetTrainees(ctx context.Context, limit, offset int) ([]*model.Trainee, error)
	UpdateTrainee(ctx context.Context, trainee *model.Trainee) error
	DeleteTrainee(ctx context.Context, id int64) error
	
	// Search and filtering operations
	SearchTrainees(ctx context.Context, query string, limit, offset int) ([]*model.Trainee, error)
	GetTraineesByEmail(ctx context.Context, email string) (*model.Trainee, error)
	GetTraineesByPhone(ctx context.Context, phone string) (*model.Trainee, error)
	GetActiveTrainees(ctx context.Context, limit, offset int) ([]*model.Trainee, error)
	GetTraineesByMembershipType(ctx context.Context, membershipType model.MembershipType, limit, offset int) ([]*model.Trainee, error)
	GetTraineesByFitnessLevel(ctx context.Context, fitnessLevel model.FitnessLevel, limit, offset int) ([]*model.Trainee, error)
	
	// Membership operations
	GetExpiringMemberships(ctx context.Context, days int, limit, offset int) ([]*model.Trainee, error)
	GetExpiredMemberships(ctx context.Context, limit, offset int) ([]*model.Trainee, error)
	UpdateMembershipStatus(ctx context.Context, id int64, isActive bool) error
}
