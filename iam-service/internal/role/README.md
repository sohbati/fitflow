# Role Management System

This document describes the role management system added to the IAM service, providing comprehensive role management capabilities.

## Overview

The role management system allows administrators to create, read, update, and delete roles that can be assigned to users. Each role has a unique name and description.

## Role Model

### Database Schema

```sql
CREATE TABLE roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    role VARCHAR UNIQUE NOT NULL,
    description VARCHAR NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
```

### Role Structure

```go
type Role struct {
    ID          uuid.UUID `json:"id"`
    Role        string    `json:"role"`        // Unique role name
    Description string    `json:"description"` // Role description
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

## API Endpoints

All role management endpoints are protected and require JWT authentication.

### Base URL
```
/api/v1/roles
```

### 1. Create Role

**POST** `/api/v1/roles`

Creates a new role.

**Request Body:**
```json
{
  "role": "admin",
  "description": "Administrator with full access"
}
```

**Response (201 Created):**
```json
{
  "message": "Role created successfully",
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "role": "admin",
    "description": "Administrator with full access",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

**Error Responses:**
- `400 Bad Request` - Validation error
- `409 Conflict` - Role already exists
- `500 Internal Server Error` - Server error

### 2. Get All Roles

**GET** `/api/v1/roles`

Retrieves all available roles.

**Response (200 OK):**
```json
{
  "message": "Roles retrieved successfully",
  "data": [
    {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "role": "admin",
      "description": "Administrator with full access",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    },
    {
      "id": "456e7890-e89b-12d3-a456-426614174001",
      "role": "user",
      "description": "Regular user with basic access",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

### 3. Get Role by ID

**GET** `/api/v1/roles/{id}`

Retrieves a specific role by its ID.

**Path Parameters:**
- `id` (string, required) - Role UUID

**Response (200 OK):**
```json
{
  "message": "Role retrieved successfully",
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "role": "admin",
    "description": "Administrator with full access",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

**Error Responses:**
- `400 Bad Request` - Invalid ID format
- `404 Not Found` - Role not found
- `500 Internal Server Error` - Server error

### 4. Update Role

**PUT** `/api/v1/roles/{id}`

Updates an existing role.

**Path Parameters:**
- `id` (string, required) - Role UUID

**Request Body:**
```json
{
  "role": "super_admin",
  "description": "Super administrator with enhanced privileges"
}
```

**Response (200 OK):**
```json
{
  "message": "Role updated successfully",
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "role": "super_admin",
    "description": "Super administrator with enhanced privileges",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
  }
}
```

**Error Responses:**
- `400 Bad Request` - Validation error or invalid ID
- `404 Not Found` - Role not found
- `409 Conflict` - Role name already exists
- `500 Internal Server Error` - Server error

### 5. Delete Role

**DELETE** `/api/v1/roles/{id}`

Deletes a role by its ID.

**Path Parameters:**
- `id` (string, required) - Role UUID

**Response (200 OK):**
```json
{
  "message": "Role deleted successfully",
  "data": null
}
```

**Error Responses:**
- `400 Bad Request` - Invalid ID format
- `404 Not Found` - Role not found
- `500 Internal Server Error` - Server error

## Usage Examples

### 1. Create a new role
```bash
curl -X POST http://localhost:8091/api/v1/roles \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-jwt-token>" \
  -d '{
    "role": "moderator",
    "description": "Moderator with content management access"
  }'
```

### 2. Get all roles
```bash
curl -X GET http://localhost:8091/api/v1/roles \
  -H "Authorization: Bearer <your-jwt-token>"
```

### 3. Get specific role
```bash
curl -X GET http://localhost:8091/api/v1/roles/123e4567-e89b-12d3-a456-426614174000 \
  -H "Authorization: Bearer <your-jwt-token>"
```

### 4. Update role
```bash
curl -X PUT http://localhost:8091/api/v1/roles/123e4567-e89b-12d3-a456-426614174000 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-jwt-token>" \
  -d '{
    "description": "Updated moderator description"
  }'
```

### 5. Delete role
```bash
curl -X DELETE http://localhost:8091/api/v1/roles/123e4567-e89b-12d3-a456-426614174000 \
  -H "Authorization: Bearer <your-jwt-token>"
```

## Architecture

### Repository Pattern
- **Repository Interface** - Defines data access methods
- **Repository Implementation** - GORM-based implementation
- **Service Layer** - Business logic and validation
- **Handler Layer** - HTTP request/response handling

### Components

1. **Model** (`internal/role/model.go`)
   - Role entity definition
   - GORM tags for database mapping

2. **Repository** (`internal/role/repository.go`)
   - Data access interface and implementation
   - CRUD operations

3. **Service** (`internal/role/service.go`)
   - Business logic
   - Validation and error handling

4. **Handler** (`internal/role/handler.go`)
   - HTTP request handling
   - Swagger documentation annotations

## Error Handling

The role management system provides comprehensive error handling:

### Error Types
- **Validation Errors** - Invalid request data
- **Not Found Errors** - Role doesn't exist
- **Conflict Errors** - Role name already exists
- **Server Errors** - Internal server issues

### Error Response Format
```json
{
  "error": "error_code",
  "message": "Human readable error message"
}
```

## Security

- **JWT Authentication** - All endpoints require valid JWT token
- **Authorization Header** - Must include `Bearer <token>`
- **Protected Routes** - All role management endpoints are protected

## Database Integration

- **Auto Migration** - Roles table is automatically created
- **UUID Primary Keys** - Unique identifiers for roles
- **Timestamps** - Automatic created_at and updated_at tracking
- **Unique Constraints** - Role names must be unique

## Swagger Documentation

The role management endpoints are fully documented in Swagger:
- Visit `http://localhost:8091/swagger/index.html`
- All endpoints include request/response schemas
- Interactive testing available

## Future Enhancements

Potential future improvements:
1. **Role Permissions** - Add permission system to roles
2. **Role Hierarchy** - Implement role inheritance
3. **Bulk Operations** - Add bulk create/update/delete
4. **Role Assignment** - Direct role assignment to users
5. **Audit Logging** - Track role changes
6. **Soft Delete** - Implement soft delete for roles

This role management system provides a solid foundation for managing user roles in the IAM service.
