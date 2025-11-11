# Trainee Entity Documentation

This document describes the `trainee` entity in the FitFlow Business service, which represents gym members/trainees.

## Overview

The `trainee` entity manages gym members and their information, including personal details, fitness metrics, membership status, and emergency contacts.

## Database Schema

### Table: `trainees`

| Column | Type | Description |
|--------|------|-------------|
| `id` | BIGSERIAL | Primary key |
| `name` | VARCHAR(150) | Full name (required) |
| `email` | VARCHAR(100) | Email address (unique) |
| `phone_number` | VARCHAR(30) | Phone number |
| `date_of_birth` | DATE | Date of birth for age calculation |
| `gender` | VARCHAR(10) | Gender: male, female, other |
| `height_cm` | INTEGER | Height in centimeters |
| `weight_kg` | DECIMAL(5,2) | Weight in kilograms |
| `fitness_level` | VARCHAR(20) | beginner, intermediate, advanced |
| `goals` | TEXT | Fitness goals and objectives |
| `medical_conditions` | TEXT | Medical conditions or restrictions |
| `emergency_contact_name` | VARCHAR(150) | Emergency contact person |
| `emergency_contact_phone` | VARCHAR(30) | Emergency contact phone |
| `emergency_contact_relation` | VARCHAR(50) | Relationship to emergency contact |
| `membership_type` | VARCHAR(50) | basic, premium, vip |
| `membership_start_date` | DATE | Membership start date |
| `membership_end_date` | DATE | Membership end date |
| `is_active` | BOOLEAN | Whether trainee is currently active |
| `profile_image_url` | VARCHAR(500) | URL to profile image |
| `created_at` | TIMESTAMP | Record creation timestamp |
| `updated_at` | TIMESTAMP | Record update timestamp |

## Model Structure

### Go Model (`internal/model/trainee.go`)

```go
type Trainee struct {
    ID                      int64           `json:"id"`
    Name                    string          `json:"name"`
    Email                   *string         `json:"email"`
    PhoneNumber             *string         `json:"phone_number"`
    DateOfBirth             *time.Time      `json:"date_of_birth"`
    Gender                  *Gender         `json:"gender"`
    HeightCm                *int            `json:"height_cm"`
    WeightKg                *float64        `json:"weight_kg"`
    FitnessLevel            FitnessLevel    `json:"fitness_level"`
    Goals                   *string         `json:"goals"`
    MedicalConditions       *string         `json:"medical_conditions"`
    EmergencyContactName    *string         `json:"emergency_contact_name"`
    EmergencyContactPhone   *string         `json:"emergency_contact_phone"`
    EmergencyContactRelation *string        `json:"emergency_contact_relation"`
    MembershipType          MembershipType  `json:"membership_type"`
    MembershipStartDate     *time.Time      `json:"membership_start_date"`
    MembershipEndDate       *time.Time      `json:"membership_end_date"`
    IsActive                bool            `json:"is_active"`
    ProfileImageURL         *string         `json:"profile_image_url"`
    CreatedAt               time.Time       `json:"created_at"`
    UpdatedAt               time.Time       `json:"updated_at"`
}
```

### Custom Types

#### Gender
```go
type Gender string
const (
    GenderMale   Gender = "male"
    GenderFemale Gender = "female"
    GenderOther  Gender = "other"
)
```

#### FitnessLevel
```go
type FitnessLevel string
const (
    FitnessLevelBeginner     FitnessLevel = "beginner"
    FitnessLevelIntermediate FitnessLevel = "intermediate"
    FitnessLevelAdvanced     FitnessLevel = "advanced"
)
```

#### MembershipType
```go
type MembershipType string
const (
    MembershipTypeBasic   MembershipType = "basic"
    MembershipTypePremium MembershipType = "premium"
    MembershipTypeVIP     MembershipType = "vip"
)
```

## API Endpoints

### Basic CRUD Operations

#### Create Trainee
- **POST** `/api/v1/trainees`
- **Body**: Trainee object
- **Response**: Created trainee with ID

#### Get Trainee
- **GET** `/api/v1/trainees/:id`
- **Response**: Trainee object

#### Update Trainee
- **PUT** `/api/v1/trainees/:id`
- **Body**: Updated trainee object
- **Response**: Updated trainee

#### Delete Trainee
- **DELETE** `/api/v1/trainees/:id`
- **Response**: 204 No Content

#### List Trainees
- **GET** `/api/v1/trainees?limit=10&offset=0`
- **Response**: Array of trainees

### Search and Filter Operations

#### Search Trainees
- **GET** `/api/v1/trainees/search?q=query&limit=10&offset=0`
- **Response**: Array of matching trainees

#### Get Active Trainees
- **GET** `/api/v1/trainees/active?limit=10&offset=0`
- **Response**: Array of active trainees

#### Get Trainees by Email
- **GET** `/api/v1/trainees/email/:email`
- **Response**: Trainee object

#### Get Trainees by Phone
- **GET** `/api/v1/trainees/phone/:phone`
- **Response**: Trainee object

#### Get Trainees by Membership Type
- **GET** `/api/v1/trainees/membership/:type?limit=10&offset=0`
- **Response**: Array of trainees

#### Get Trainees by Fitness Level
- **GET** `/api/v1/trainees/fitness/:level?limit=10&offset=0`
- **Response**: Array of trainees

### Membership Management

#### Get Expiring Memberships
- **GET** `/api/v1/trainees/expiring?days=30&limit=10&offset=0`
- **Response**: Array of trainees with expiring memberships

#### Get Expired Memberships
- **GET** `/api/v1/trainees/expired?limit=10&offset=0`
- **Response**: Array of trainees with expired memberships

#### Update Membership Status
- **PATCH** `/api/v1/trainees/:id/status`
- **Body**: `{"is_active": true/false}`
- **Response**: Success message

### Health and Fitness Calculations

#### Calculate Age
- **GET** `/api/v1/trainees/:id/age`
- **Response**: `{"age": 25}`

#### Calculate BMI
- **GET** `/api/v1/trainees/:id/bmi`
- **Response**: `{"bmi": 22.5}`

#### Validate Membership
- **GET** `/api/v1/trainees/:id/membership/validate`
- **Response**: `{"is_valid": true}`

## Business Logic

### Validation Rules

1. **Name**: Required field
2. **Email**: Must be unique if provided
3. **Phone**: Must be unique if provided
4. **Fitness Level**: Defaults to "beginner"
5. **Membership Type**: Defaults to "basic"

### Calculated Fields

#### Age Calculation
- Based on `date_of_birth` field
- Returns current age in years

#### BMI Calculation
- Formula: `weight(kg) / height(m)^2`
- Requires both height and weight to be provided

#### Membership Validation
- Checks if `is_active` is true
- Checks if `membership_end_date` is in the future (if provided)

## Usage Examples

### Creating a Trainee
```json
POST /api/v1/trainees
{
    "name": "John Doe",
    "email": "john@example.com",
    "phone_number": "+1234567890",
    "date_of_birth": "1990-01-01",
    "gender": "male",
    "height_cm": 180,
    "weight_kg": 75.5,
    "fitness_level": "intermediate",
    "goals": "Build muscle and improve endurance",
    "membership_type": "premium",
    "membership_start_date": "2024-01-01",
    "membership_end_date": "2024-12-31"
}
```

### Searching Trainees
```
GET /api/v1/trainees/search?q=john&limit=10&offset=0
```

### Getting Expiring Memberships
```
GET /api/v1/trainees/expiring?days=30&limit=10&offset=0
```

## Database Indexes

The following indexes are created for optimal performance:

- `idx_trainees_email` - On email field
- `idx_trainees_phone` - On phone_number field
- `idx_trainees_membership_type` - On membership_type field
- `idx_trainees_is_active` - On is_active field
- `idx_trainees_membership_dates` - On membership_start_date and membership_end_date
- `idx_trainees_fitness_level` - On fitness_level field

## File Structure

```
fitflow-business/
├── database/postgres/
│   └── 6_trainees.sql                    # Database schema
├── internal/
│   ├── model/
│   │   └── trainee.go                    # Go model
│   ├── repository/
│   │   ├── trainee_repository.go         # Repository interface
│   │   └── impl/
│   │       └── trainee_repository_impl.go # Repository implementation
│   ├── service/
│   │   ├── trainee_service.go            # Service interface
│   │   └── trainee_service_impl.go       # Service implementation
│   ├── handler/
│   │   └── trainee_handler_gin.go        # HTTP handlers
│   └── router/
│       └── router.go                     # Route definitions
└── docker-compose.yml                    # Docker configuration
```

This trainee entity provides comprehensive management of gym members with features for membership tracking, health metrics, and flexible search capabilities.
