# Person-Based Architecture Documentation

This document describes the new person-based architecture in the FitFlow Business service, where all user roles (gym owners, trainers, trainees) inherit from a common `person` entity.

## Architecture Overview

The system now follows an inheritance-based design where:

1. **Person** - Base entity containing common personal information
2. **Gym Owner** - Inherits from Person, adds gym-specific information
3. **Trainer** - Inherits from Person, adds trainer-specific information  
4. **Trainee** - Inherits from Person, adds trainee-specific information

Each person is linked to a user in the IAM system via `user_id`.

## Database Schema

### Person Table (`persons`)

| Column | Type | Description |
|--------|------|-------------|
| `id` | BIGSERIAL | Primary key |
| `user_id` | UUID | Reference to users table in IAM system |
| `first_name` | VARCHAR(100) | First name (required) |
| `last_name` | VARCHAR(100) | Last name (required) |
| `email` | VARCHAR(100) | Email address (unique) |
| `phone_number` | VARCHAR(30) | Phone number |
| `date_of_birth` | DATE | Date of birth for age calculation |
| `gender` | VARCHAR(10) | Gender: male, female, other |
| `profile_image_url` | VARCHAR(500) | URL to profile image |
| `address` | VARCHAR(255) | Street address |
| `city` | VARCHAR(100) | City name |
| `province` | VARCHAR(100) | Province/state name |
| `country` | VARCHAR(100) | Country name |
| `postal_code` | VARCHAR(20) | Postal/ZIP code |
| `emergency_contact_name` | VARCHAR(150) | Emergency contact person |
| `emergency_contact_phone` | VARCHAR(30) | Emergency contact phone |
| `emergency_contact_relation` | VARCHAR(50) | Relationship to emergency contact |
| `is_active` | BOOLEAN | Whether person is currently active |
| `created_at` | TIMESTAMP | Record creation timestamp |
| `updated_at` | TIMESTAMP | Record update timestamp |

### Role-Specific Tables

#### Gym Owners (`gym_owners`)
- `id` - Primary key
- `person_id` - Reference to persons table
- `gym_id` - Reference to gyms table
- `brief_bio` - Short description about the owner

#### Trainers (`trainers`)
- `id` - Primary key
- `person_id` - Reference to persons table
- `is_registered` - Whether trainer is registered with a gym

#### Trainees (`trainees`)
- `id` - Primary key
- `person_id` - Reference to persons table
- `height_cm` - Height in centimeters
- `weight_kg` - Weight in kilograms
- `fitness_level` - beginner, intermediate, advanced
- `goals` - Fitness goals and objectives
- `medical_conditions` - Medical conditions or restrictions
- `membership_type` - basic, premium, vip
- `membership_start_date` - Membership start date
- `membership_end_date` - Membership end date
- `is_active` - Whether trainee is currently active

## Go Models

### Person Model

```go
type Person struct {
    ID                      int64      `json:"id"`
    UserID                  uuid.UUID  `json:"user_id"`
    FirstName               string     `json:"first_name"`
    LastName                string     `json:"last_name"`
    Email                   *string    `json:"email"`
    PhoneNumber             *string    `json:"phone_number"`
    DateOfBirth             *time.Time `json:"date_of_birth"`
    Gender                  *Gender    `json:"gender"`
    ProfileImageURL         *string    `json:"profile_image_url"`
    Address                 *string    `json:"address"`
    City                    *string    `json:"city"`
    Province                *string    `json:"province"`
    Country                 *string    `json:"country"`
    PostalCode              *string    `json:"postal_code"`
    EmergencyContactName    *string    `json:"emergency_contact_name"`
    EmergencyContactPhone   *string    `json:"emergency_contact_phone"`
    EmergencyContactRelation *string   `json:"emergency_contact_relation"`
    IsActive                bool       `json:"is_active"`
    CreatedAt               time.Time  `json:"created_at"`
    UpdatedAt               time.Time  `json:"updated_at"`
}
```

### Role Models

#### Gym Owner Model
```go
type GymOwner struct {
    ID        int64     `json:"id"`
    PersonID  int64     `json:"person_id"`
    GymID     int64     `json:"gym_id"`
    BriefBio  *string   `json:"brief_bio"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    
    // Relationships
    Person Person `json:"person,omitempty" gorm:"foreignKey:PersonID"`
    Gym    Gym    `json:"gym,omitempty" gorm:"foreignKey:GymID"`
}
```

#### Trainer Model
```go
type Trainer struct {
    ID           int64     `json:"id"`
    PersonID     int64     `json:"person_id"`
    IsRegistered bool      `json:"is_registered"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
    
    // Relationships
    Person Person `json:"person,omitempty" gorm:"foreignKey:PersonID"`
    Gyms   []Gym  `json:"gyms,omitempty" gorm:"many2many:gym_trainers;"`
}
```

#### Trainee Model
```go
type Trainee struct {
    ID                  int64           `json:"id"`
    PersonID            int64           `json:"person_id"`
    HeightCm            *int            `json:"height_cm"`
    WeightKg            *float64        `json:"weight_kg"`
    FitnessLevel        FitnessLevel    `json:"fitness_level"`
    Goals               *string         `json:"goals"`
    MedicalConditions   *string         `json:"medical_conditions"`
    MembershipType      MembershipType  `json:"membership_type"`
    MembershipStartDate *time.Time      `json:"membership_start_date"`
    MembershipEndDate   *time.Time      `json:"membership_end_date"`
    IsActive            bool            `json:"is_active"`
    CreatedAt           time.Time       `json:"created_at"`
    UpdatedAt           time.Time       `json:"updated_at"`
    
    // Relationships
    Person Person `json:"person,omitempty" gorm:"foreignKey:PersonID"`
}
```

## Benefits of This Architecture

### 1. **Code Reusability**
- Common personal information is stored once in the `persons` table
- Shared functionality (age calculation, full name) is implemented once
- Reduces code duplication across role-specific implementations

### 2. **Data Consistency**
- Single source of truth for personal information
- Email and phone uniqueness is enforced at the person level
- Consistent data validation across all roles

### 3. **User Management Integration**
- Direct link to IAM system via `user_id`
- Single person can have multiple roles (e.g., trainer who is also a trainee)
- Centralized user profile management

### 4. **Scalability**
- Easy to add new roles by creating new tables that reference `persons`
- Role-specific data is separated from common data
- Flexible querying capabilities

### 5. **Maintainability**
- Changes to common fields only need to be made in one place
- Clear separation of concerns between common and role-specific data
- Easier to understand and maintain the data model

## Usage Examples

### Creating a Person with Multiple Roles

```go
// 1. Create person
person := &model.Person{
    UserID:    userID,
    FirstName: "John",
    LastName:  "Doe",
    Email:     "john@example.com",
    PhoneNumber: "+1234567890",
    // ... other fields
}

// 2. Create as trainer
trainer := &model.Trainer{
    PersonID:     person.ID,
    IsRegistered: true,
}

// 3. Create as trainee
trainee := &model.Trainee{
    PersonID:       person.ID,
    FitnessLevel:   model.FitnessLevelBeginner,
    MembershipType: model.MembershipTypeBasic,
    // ... other fields
}
```

### Querying with Person Information

```go
// Get trainer with person information
var trainer model.Trainer
db.Preload("Person").First(&trainer, trainerID)

// Access person information
fullName := trainer.Person.GetFullName()
age, _ := trainer.Person.GetAge()
```

## Migration Strategy

### From Old Schema to New Schema

1. **Create persons table** with all common fields
2. **Migrate existing data** from role tables to persons table
3. **Update role tables** to reference person_id instead of storing personal info
4. **Update application code** to use new models and relationships
5. **Test thoroughly** to ensure data integrity

### Data Migration Example

```sql
-- Migrate gym owners to persons
INSERT INTO persons (user_id, first_name, last_name, email, phone_number, created_at, updated_at)
SELECT 
    gen_random_uuid(), -- Generate temporary user_id
    split_part(name, ' ', 1) as first_name,
    split_part(name, ' ', 2) as last_name,
    email,
    phone_number,
    created_at,
    updated_at
FROM gym_owners;

-- Update gym_owners to reference persons
UPDATE gym_owners 
SET person_id = persons.id 
FROM persons 
WHERE persons.email = gym_owners.email;
```

## File Structure

```
fitflow-business/
├── database/postgres/
│   ├── 7_persons.sql                    # Person schema
│   ├── 3_gym_owners.sql                 # Updated gym owners schema
│   ├── 4_trainers.sql                   # Updated trainers schema
│   └── 6_trainees.sql                   # Updated trainees schema
├── internal/
│   ├── model/
│   │   ├── person.go                    # Person model
│   │   ├── gym_owner.go                 # Updated gym owner model
│   │   ├── trainer.go                   # Updated trainer model
│   │   └── trainee.go                   # Updated trainee model
│   ├── repository/
│   │   ├── person_repository.go         # Person repository interface
│   │   └── impl/
│   │       └── person_repository_impl.go # Person repository implementation
│   └── service/
│       ├── person_service.go            # Person service interface
│       └── person_service_impl.go       # Person service implementation
└── docker-compose.yml                   # Updated Docker configuration
```

This person-based architecture provides a solid foundation for user management while maintaining flexibility for role-specific functionality.
