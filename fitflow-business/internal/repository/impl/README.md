# Repository Implementations

This directory contains the concrete implementations of the repository interfaces defined in the parent directory.

## Purpose

The `impl/` directory follows the **Interface Segregation Principle** and **Dependency Inversion Principle** by:

- Separating interface definitions from their implementations
- Making implementations easily replaceable
- Providing a clear boundary between contracts and concrete code
- Enabling better testing through interface mocking

## File Structure

### `gym_repository_impl.go`
- **Implements**: `repository.GymRepository`
- **Features**: GORM-based implementation with context support
- **Operations**: Full CRUD, search, verification filtering

### `gym_location_repository_impl.go`
- **Implements**: `repository.GymLocationRepository`
- **Features**: Location-specific operations with gym relationships
- **Operations**: CRUD operations for gym locations

### `gym_owner_repository_impl.go`
- **Implements**: `repository.GymOwnerRepository`
- **Features**: Owner-specific operations with gym relationships
- **Operations**: CRUD operations for gym owners

### `trainer_repository_impl.go`
- **Implements**: `repository.TrainerRepository`
- **Features**: Trainer-specific operations with gym relationships
- **Operations**: CRUD operations, registered trainer filtering

### `gym_trainer_repository_impl.go`
- **Implements**: `repository.GymTrainerRepository`
- **Features**: Many-to-many relationship management
- **Operations**: Add/remove relationships, complex joins

## Implementation Details

### GORM Integration
All implementations use GORM for database operations:
- Context support for cancellation and timeouts
- Preloading relationships to avoid N+1 queries
- Efficient pagination and filtering
- Complex joins for relationship queries

### Error Handling
- Proper error propagation from database layer
- Context-aware error handling
- GORM error translation

### Performance Optimizations
- Preloading related entities
- Efficient query construction
- Pagination support
- Search optimization with ILIKE

## Usage

```go
import (
    "fitflow-business/internal/repository"
    "fitflow-business/internal/repository/impl"
)

// Create implementations
gymRepo := impl.NewGymRepository(db)
locationRepo := impl.NewGymLocationRepository(db)
// ... other repositories

// Use in services
service := service.NewGymService(gymRepo, locationRepo, ...)
```

## Testing

These implementations can be easily tested by:
1. Using in-memory databases (SQLite)
2. Mocking the interfaces for unit tests
3. Integration testing with test databases

## Future Enhancements

- Add caching layer
- Implement query optimization
- Add metrics and monitoring
- Support for different database backends
