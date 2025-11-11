# Gym Entities and API Documentation

## Database Schema

The gym-related entities are implemented with the following tables:

### 1. `gyms` Table
- **Primary Key**: `id` (BIGSERIAL)
- **Fields**: name, description, phone_number, email, website_url, is_verified, facilities (JSONB), images (JSONB)
- **Features**: JSONB fields for flexible data storage, automatic timestamps

### 2. `gym_locations` Table
- **Primary Key**: `id` (BIGSERIAL)
- **Foreign Key**: `gym_id` references `gyms(id)`
- **Fields**: location_type, address, city, province, country, postal_code, latitude, longitude
- **Features**: Support for multiple locations per gym, geographic coordinates

### 3. `gym_owners` Table
- **Primary Key**: `id` (BIGSERIAL)
- **Foreign Key**: `gym_id` references `gyms(id)`
- **Fields**: name, phone_number, email, brief_bio
- **Features**: Support for multiple owners per gym

### 4. `trainers` Table
- **Primary Key**: `id` (BIGSERIAL)
- **Fields**: name, phone_number, email, is_registered
- **Features**: Independent trainers or gym-registered trainers

### 5. `gym_trainers` Table (Junction Table)
- **Primary Key**: `id` (BIGSERIAL)
- **Foreign Keys**: `gym_id` references `gyms(id)`, `trainer_id` references `trainers(id)`
- **Features**: Many-to-many relationship between gyms and trainers

## API Endpoints

### Gym Management
- `GET /api/v1/gyms` - List all gyms (with pagination)
- `GET /api/v1/gyms/search?q={query}` - Search gyms by name/description
- `GET /api/v1/gyms/verified` - Get only verified gyms
- `GET /api/v1/gyms/{id}` - Get gym by ID
- `POST /api/v1/gyms` - Create new gym
- `PUT /api/v1/gyms/{id}` - Update gym
- `DELETE /api/v1/gyms/{id}` - Delete gym

### Gym Locations
- `GET /api/v1/gyms/{id}/locations` - Get gym locations
- `POST /api/v1/gyms/{id}/locations` - Add location to gym

### Gym Owners
- `GET /api/v1/gyms/{id}/owners` - Get gym owners
- `POST /api/v1/gyms/{id}/owners` - Add owner to gym

### Trainers
- `GET /api/v1/trainers` - List all trainers (with pagination)
- `GET /api/v1/trainers/{id}` - Get trainer by ID
- `POST /api/v1/trainers` - Create new trainer

### Gym-Trainer Relationships
- `GET /api/v1/gyms/{id}/trainers` - Get trainers for a gym
- `POST /api/v1/gyms/{id}/trainers/{trainer_id}` - Add trainer to gym

## JSONB Fields Usage

### Facilities Example
```json
{
  "wifi": true,
  "pool": false,
  "sauna": true,
  "parking": true,
  "locker_rooms": true,
  "shower_facilities": true
}
```

### Images Example
```json
[
  {
    "url": "https://example.com/logo.png",
    "type": "logo",
    "is_primary": true
  },
  {
    "url": "https://example.com/interior1.png",
    "type": "interior",
    "is_primary": false
  },
  {
    "url": "https://example.com/trainer1.png",
    "type": "trainer",
    "is_primary": false
  }
]
```

## Architecture

The implementation follows clean architecture principles:

1. **Model Layer** (`internal/model/gym.go`): Entity definitions with GORM tags
2. **Repository Layer** (`internal/repository/`): Database access abstraction
3. **Service Layer** (`internal/service/`): Business logic and validation
4. **Handler Layer** (`internal/handler/`): HTTP request/response handling
5. **Router Layer** (`internal/router/`): Route definitions and middleware

## Features

- **JSONB Support**: Flexible storage for facilities and images
- **Automatic Timestamps**: Created/updated timestamps with triggers
- **Foreign Key Constraints**: Proper referential integrity
- **Indexes**: Performance optimization for common queries
- **Validation**: Input validation in service layer
- **Pagination**: Built-in pagination support
- **Search**: Full-text search capabilities
- **Relationships**: Proper handling of gym-trainer many-to-many relationships
