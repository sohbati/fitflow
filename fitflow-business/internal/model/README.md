# FitFlow Business Model Files

This directory contains the Go model files for the FitFlow Business service, organized by entity for better maintainability and separation of concerns.

## File Structure

### `jsonb.go`
- **Purpose**: Common JSONB type for PostgreSQL compatibility
- **Contains**: `JSONB` type with `Value()` and `Scan()` methods
- **Usage**: Used by other models that need JSONB fields

### `image.go`
- **Purpose**: Image-related types for gym images
- **Contains**: `Image` struct and `Images` slice type
- **Features**: JSONB serialization for PostgreSQL storage
- **Usage**: Used by `Gym` model for storing image data

### `gym.go`
- **Purpose**: Main gym entity
- **Contains**: `Gym` struct with all gym-related fields
- **Relationships**: References to locations, owners, and trainers
- **Features**: JSONB fields for facilities and images

### `gym_location.go`
- **Purpose**: Gym location entity
- **Contains**: `GymLocation` struct for gym locations
- **Relationships**: Belongs to a gym
- **Features**: Geographic coordinates (latitude/longitude)

### `gym_owner.go`
- **Purpose**: Gym owner entity
- **Contains**: `GymOwner` struct for gym owners
- **Relationships**: Belongs to a gym
- **Features**: Owner contact information and bio

### `trainer.go`
- **Purpose**: Trainer entity
- **Contains**: `Trainer` struct for personal trainers
- **Relationships**: Many-to-many with gyms
- **Features**: Independent or gym-registered trainers

### `gym_trainer.go`
- **Purpose**: Junction table entity
- **Contains**: `GymTrainer` struct for gym-trainer relationships
- **Relationships**: Links gyms and trainers
- **Features**: Many-to-many relationship implementation

## Model Relationships

```
Gym (1) ──→ (N) GymLocation
Gym (1) ──→ (N) GymOwner
Gym (N) ←──→ (N) Trainer (via GymTrainer)
```

## Key Features

### JSONB Support
- **Facilities**: Flexible storage for gym amenities
- **Images**: Array of image objects with metadata
- **Type Safety**: Custom types with proper serialization

### GORM Integration
- **Tags**: Proper GORM tags for database mapping
- **Relationships**: Foreign keys and many-to-many relationships
- **Timestamps**: Automatic created_at and updated_at handling

### Geographic Data
- **Coordinates**: Latitude and longitude for locations
- **Address Fields**: Complete address information
- **Location Types**: Support for different location types

## Usage Examples

### Creating a Gym with Images
```go
gym := &model.Gym{
    Name: "FitFlow Gym",
    Description: "A modern fitness center",
    Facilities: model.JSONB{
        "wifi": true,
        "pool": false,
        "sauna": true,
    },
    Images: model.Images{
        {
            URL: "https://example.com/logo.png",
            Type: "logo",
            IsPrimary: true,
        },
    },
}
```

### Adding a Location
```go
location := &model.GymLocation{
    GymID: gym.ID,
    LocationType: "gym",
    Address: "123 Main St",
    City: "New York",
    Latitude: 40.7128,
    Longitude: -74.0060,
}
```

### Linking Trainer to Gym
```go
gymTrainer := &model.GymTrainer{
    GymID: gym.ID,
    TrainerID: trainer.ID,
}
```

## Benefits of Separation

1. **Modularity**: Each entity in its own file
2. **Maintainability**: Easy to modify individual models
3. **Clarity**: Clear separation of concerns
4. **Reusability**: Models can be imported independently
5. **Testing**: Easier to write focused unit tests
6. **Documentation**: Better organization for documentation

## Import Usage

```go
import "fitflow-business/internal/model"

// Use any model
gym := &model.Gym{}
location := &model.GymLocation{}
trainer := &model.Trainer{}
```

This modular approach makes the codebase more maintainable and follows Go best practices for package organization.
