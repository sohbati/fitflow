# User Model Documentation

This document describes the updated user model with email and mobile fields added to the IAM service.

## User Model Structure

### Database Schema

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR UNIQUE NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    mobile VARCHAR UNIQUE NOT NULL,
    display_name VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    country VARCHAR NOT NULL,
    role VARCHAR DEFAULT 'user',
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
```

### Go Model

```go
type User struct {
    ID          uuid.UUID `json:"id"`
    Username    string    `json:"username"`     // Unique username
    Email       string    `json:"email"`        // Unique email address
    Mobile      string    `json:"mobile"`       // Unique mobile number
    DisplayName string    `json:"display_name"` // Display name
    Password    string    `json:"-"`            // Hashed password (hidden in JSON)
    Country     string    `json:"country"`      // ISO country code
    Role        string    `json:"-"`            // Hidden from JSON responses
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

## Field Details

### Required Fields

| Field | Type | Description | Validation |
|-------|------|-------------|------------|
| `username` | string | Unique username | Required, unique |
| `email` | string | Email address | Required, unique, email format |
| `mobile` | string | Mobile number | Required, unique |
| `display_name` | string | Display name | Required |
| `password` | string | Password | Required, will be hashed |
| `country` | string | Country code | Required |

### Optional Fields

| Field | Type | Description | Default |
|-------|------|-------------|---------|
| `role` | string | User role | "user" |

### System Fields

| Field | Type | Description |
|-------|------|-------------|
| `id` | uuid.UUID | Primary key |
| `created_at` | time.Time | Creation timestamp |
| `updated_at` | time.Time | Last update timestamp |

### Hidden Fields

| Field | Type | Description | Reason for Hiding |
|-------|------|-------------|-------------------|
| `password` | string | Hashed password | Security - Never expose passwords |
| `role` | string | User role | Security - Keep role information private |

## API Changes

### Registration Endpoint

**POST** `/register` or `/auth/register`

**Request Body:**
```json
{
  "username": "john_doe",
  "email": "john.doe@example.com",
  "mobile": "+1234567890",
  "display_name": "John Doe",
  "password": "securepassword123",
  "country": "US",
  "role": "user"
}
```

**Response (201 Created):**
```json
{
  "message": "User created successfully",
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "username": "john_doe",
    "email": "john.doe@example.com",
    "mobile": "+1234567890",
    "display_name": "John Doe",
    "country": "US",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### Login Endpoint

**POST** `/login` or `/auth/login`

**Request Body:**
```json
{
  "username": "john_doe",
  "password": "securepassword123"
}
```

**Response (200 OK):**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "username": "john_doe",
    "email": "john.doe@example.com",
    "mobile": "+1234567890",
    "display_name": "John Doe",
    "country": "US",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

## Validation Rules

### Email Validation
- **Required**: Email field is mandatory
- **Format**: Must be a valid email format
- **Uniqueness**: Email must be unique across all users
- **Example**: `john.doe@example.com`

### Mobile Validation
- **Required**: Mobile field is mandatory
- **Uniqueness**: Mobile number must be unique across all users
- **Format**: Accepts any string format (e.g., `+1234567890`, `123-456-7890`)
- **Example**: `+1234567890`

### Username Validation
- **Required**: Username field is mandatory
- **Uniqueness**: Username must be unique across all users
- **Example**: `john_doe`

## Error Handling

### Registration Errors

**Email Already Exists (409 Conflict):**
```json
{
  "error": "creation_failed",
  "message": "email already exists"
}
```

**Mobile Already Exists (409 Conflict):**
```json
{
  "error": "creation_failed",
  "message": "mobile number already exists"
}
```

**Username Already Exists (409 Conflict):**
```json
{
  "error": "creation_failed",
  "message": "username already exists"
}
```

**Validation Error (400 Bad Request):**
```json
{
  "error": "validation_error",
  "message": "Key: 'RegisterRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag"
}
```

## Database Constraints

### Unique Constraints
- `username` - Must be unique
- `email` - Must be unique
- `mobile` - Must be unique

### Not Null Constraints
- All fields except `role` are required
- `role` defaults to "user" if not provided

## Repository Methods

### New Methods Added

```go
// Get user by email
func (r *repository) GetByEmail(email string) (*User, error)

// Get user by mobile
func (r *repository) GetByMobile(mobile string) (*User, error)
```

### Existing Methods
```go
// Create user
func (r *repository) Create(user *User) error

// Get user by username
func (r *repository) GetByUsername(username string) (*User, error)

// Get user by ID
func (r *repository) GetByID(id uuid.UUID) (*User, error)
```

## Service Layer Changes

### CreateUserRequest Updated

```go
type CreateUserRequest struct {
    Username    string `json:"username" binding:"required"`
    Email       string `json:"email" binding:"required,email"`
    Mobile      string `json:"mobile" binding:"required"`
    DisplayName string `json:"display_name" binding:"required"`
    Password    string `json:"password" binding:"required"`
    Country     string `json:"country" binding:"required"`
    Role        string `json:"role,omitempty"`
}
```

### Validation Logic

The service now validates uniqueness for:
1. **Username** - Checked against existing usernames
2. **Email** - Checked against existing emails
3. **Mobile** - Checked against existing mobile numbers

## Migration

### Database Migration
The user table will be automatically migrated when the application starts. The new fields (`email` and `mobile`) will be added with unique constraints.

### Existing Data
- Existing users without email/mobile will need to be updated
- Consider adding migration scripts for existing data

## Usage Examples

### 1. Register User with Email and Mobile
```bash
curl -X POST http://localhost:8091/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "jane_smith",
    "email": "jane.smith@example.com",
    "mobile": "+1987654321",
    "display_name": "Jane Smith",
    "password": "securepassword123",
    "country": "CA",
    "role": "user"
  }'
```

### 2. Login with Username
```bash
curl -X POST http://localhost:8091/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "jane_smith",
    "password": "securepassword123"
  }'
```

### 3. Get User Profile
```bash
curl -X GET http://localhost:8091/me \
  -H "Authorization: Bearer <jwt-token>"
```

## Security Considerations

### Email Security
- Email addresses are stored in plain text
- Consider email verification for production use
- Email addresses are visible in API responses

### Mobile Security
- Mobile numbers are stored in plain text
- Consider SMS verification for production use
- Mobile numbers are visible in API responses

### Password Security
- Passwords are hashed using bcrypt
- Password field is excluded from JSON responses
- Password hashing uses cost factor 12

## Future Enhancements

### Potential Improvements
1. **Email Verification** - Add email verification workflow
2. **SMS Verification** - Add mobile number verification
3. **Profile Updates** - Allow users to update email/mobile
4. **Alternative Login** - Login with email or mobile instead of username
5. **Contact Methods** - Support multiple email addresses or mobile numbers
6. **Privacy Settings** - Allow users to hide email/mobile from profile

### Database Optimizations
1. **Indexes** - Add indexes on email and mobile fields
2. **Normalization** - Consider separate contact information table
3. **Audit Trail** - Track changes to email/mobile fields

This updated user model provides a more comprehensive user profile with email and mobile contact information while maintaining data integrity through unique constraints and proper validation.
