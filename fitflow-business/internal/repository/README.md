# FitFlow Business Repository Files

This directory contains the repository files for the FitFlow Business service, organized by entity for better maintainability and separation of concerns.

## File Structure

### Interface Files (Root Level)

#### `gym_repository.go`
- **Purpose**: Interface for gym-related database operations
- **Contains**: `GymRepository` interface with CRUD operations
- **Operations**: Create, Read, Update, Delete, Search, GetVerified

#### `gym_location_repository.go`
- **Purpose**: Interface for gym location database operations
- **Contains**: `GymLocationRepository` interface
- **Operations**: CRUD operations for gym locations

#### `gym_owner_repository.go`
- **Purpose**: Interface for gym owner database operations
- **Contains**: `GymOwnerRepository` interface
- **Operations**: CRUD operations for gym owners

#### `trainer_repository.go`
- **Purpose**: Interface for trainer database operations
- **Contains**: `TrainerRepository` interface
- **Operations**: CRUD operations for trainers, GetRegistered

#### `gym_trainer_repository.go`
- **Purpose**: Interface for gym-trainer relationship operations
- **Contains**: `GymTrainerRepository` interface
- **Operations**: Add/Remove relationships, Get trainers/gyms

### Implementation Files (`impl/` Directory)

#### `impl/gym_repository_impl.go`
- **Purpose**: Implementation of gym repository interface
- **Contains**: `gymRepository` struct with GORM implementation
- **Features**: Context support, Preloading relationships, Search functionality

#### `impl/gym_location_repository_impl.go`
- **Purpose**: Implementation of gym location repository interface
- **Contains**: `gymLocationRepository` struct with GORM implementation
- **Features**: Context support, Gym-specific queries

#### `impl/gym_owner_repository_impl.go`
- **Purpose**: Implementation of gym owner repository interface
- **Contains**: `gymOwnerRepository` struct with GORM implementation
- **Features**: Context support, Gym-specific queries

#### `impl/trainer_repository_impl.go`
- **Purpose**: Implementation of trainer repository interface
- **Contains**: `trainerRepository` struct with GORM implementation
- **Features**: Context support, Preloading gyms, Registered trainer filtering

#### `impl/gym_trainer_repository_impl.go`
- **Purpose**: Implementation of gym-trainer relationship repository interface
- **Contains**: `gymTrainerRepository` struct with GORM implementation
- **Features**: Context support, Junction table operations, Complex joins

## Repository Pattern Benefits

### Separation of Concerns
- Each entity has its own repository interface and implementation
- Clear boundaries between different data operations
- Easy to modify individual repositories without affecting others

### Testability
- Each repository can be mocked independently
- Unit tests can focus on specific entity operations
- Integration tests can test individual repositories

### Maintainability
- Changes to one entity don't affect others
- Easy to add new operations to specific repositories
- Clear responsibility for each repository

## Usage Examples

### Creating Repositories
```go
import (
    "fitflow-business/internal/repository"
    "fitflow-business/internal/repository/impl"
)

// Individual repositories
gymRepo := impl.NewGymRepository(db)
gymLocationRepo := impl.NewGymLocationRepository(db)
gymOwnerRepo := impl.NewGymOwnerRepository(db)
trainerRepo := impl.NewTrainerRepository(db)
gymTrainerRepo := impl.NewGymTrainerRepository(db)

// Service with all repositories
gymService := service.NewGymService(
    gymRepo, 
    gymLocationRepo, 
    gymOwnerRepo, 
    trainerRepo, 
    gymTrainerRepo,
)
```

### Using Individual Repositories
```go
// Direct repository usage
gym, err := gymRepo.GetGymByID(ctx, 1)
locations, err := gymLocationRepo.GetGymLocations(ctx, gym.ID)
trainers, err := gymTrainerRepo.GetGymTrainers(ctx, gym.ID)
```

## Key Features

### Context Support
- All operations support context for cancellation and timeouts
- Proper context propagation through the application layers

### GORM Integration
- Uses GORM for database operations
- Supports preloading relationships
- Handles complex queries and joins

### Error Handling
- Proper error propagation from database layer
- Context-aware error handling

### Performance Optimizations
- Preloading relationships to avoid N+1 queries
- Efficient pagination support
- Optimized search queries with ILIKE

## Repository Relationships

```
GymRepository ──→ Gym operations
GymLocationRepository ──→ Location operations (depends on Gym)
GymOwnerRepository ──→ Owner operations (depends on Gym)
TrainerRepository ──→ Trainer operations (independent)
GymTrainerRepository ──→ Relationship operations (depends on both Gym and Trainer)
```

## Service Layer Integration

The service layer uses dependency injection to combine all repositories:

```go
type gymService struct {
    gymRepo           repository.GymRepository
    gymLocationRepo   repository.GymLocationRepository
    gymOwnerRepo      repository.GymOwnerRepository
    trainerRepo       repository.TrainerRepository
    gymTrainerRepo    repository.GymTrainerRepository
}
```

This modular approach makes the codebase more maintainable, testable, and follows Go best practices for clean architecture.
