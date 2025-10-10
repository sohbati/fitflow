# IAM Service

A Go microservice for Identity and Access Management with JWT-based authentication, multi-database support, and Swagger documentation.

## Features

- 🔐 JWT-based authentication
- 🔑 Google OAuth 2.0 integration
- 🗄️ Multi-database support (MySQL default, PostgreSQL supported)
- 📚 Swagger/OpenAPI documentation
- 🔒 Password hashing with bcrypt
- 🛡️ Custom claims support in JWT tokens
- 🌐 RESTful HTTP API
- ⚙️ Environment-based configuration
- 👥 Role management system
- 🔧 Clean architecture with separation of concerns

## Project Structure

```
iam-service/
├── cmd/
│   └── main.go              # Alternative entry point
├── database/                 # Database initialization scripts
│   ├── mysql/               # MySQL-specific scripts
│   ├── postgres/            # PostgreSQL-specific scripts
│   └── README.md            # Database documentation
├── internal/
│   ├── auth/                # Authentication handlers
│   ├── user/                # User entity, repository, service
│   ├── role/                # Role entity, repository, service
│   ├── middleware/           # JWT verification middleware
│   ├── router/              # HTTP routing configuration
│   ├── config/              # Configuration management
│   └── db/                  # Database connection and migrations
├── pkg/
│   ├── jwt/                 # JWT utilities
│   └── crypto/              # Password hashing utilities
├── docs/                    # Swagger generated files
├── main.go                  # Main application entry point
├── docker-compose.yml       # Docker Compose configuration
├── go.mod
├── config.example           # Environment configuration example
└── README.md
```

## Prerequisites

- Go 1.21 or later
- PostgreSQL 12 or later
- Swag CLI (for generating Swagger docs)

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd iam-service
```

2. Install dependencies:
```bash
go mod tidy
```

3. Install Swag CLI (if not already installed):
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

4. Generate Swagger documentation:
```bash
swag init
```

## Configuration

Create a `.env` file in the root directory with the following variables:

```env
JWT_SECRET=mysecret
DATABASE_URL=iam_user:password@tcp(localhost:3306)/iam_db?charset=utf8mb4&parseTime=True&loc=Local
DATABASE_TYPE=mysql
TOKEN_EXP_MINUTES=15
PORT=8091
```

### Database Configuration

#### MySQL (Default)
```env
DATABASE_URL=iam_user:password@tcp(localhost:3306)/iam_db?charset=utf8mb4&parseTime=True&loc=Local
DATABASE_TYPE=mysql
```

#### PostgreSQL
```env
DATABASE_URL=postgres://iam_user:password@localhost:5432/iam_db?sslmode=disable
DATABASE_TYPE=postgres
```

Or copy from the example:
```bash
cp config.example .env
```

## Quick Start

### Using Docker Compose (Recommended)

1. Start MySQL database:
```bash
make docker-up
```

2. Run the application:
```bash
make run
```

### Using PostgreSQL

1. Start PostgreSQL database:
```bash
make docker-up-postgres
```

2. Update your `.env` file:
```env
DATABASE_URL=postgres://iam_user:password@localhost:5432/iam_db?sslmode=disable
DATABASE_TYPE=postgres
```

3. Run the application:
```bash
make run
```

## Database Setup

### MySQL (Default)

1. Create a MySQL database:
```sql
CREATE DATABASE iam_db;
CREATE USER 'iam_user'@'localhost' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON iam_db.* TO 'iam_user'@'localhost';
FLUSH PRIVILEGES;
```

2. Or use Docker Compose:
```bash
make docker-up
```

### PostgreSQL

1. Create a PostgreSQL database:
```sql
CREATE DATABASE iam_db;
CREATE USER iam_user WITH PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE iam_db TO iam_user;
```

2. Or use Docker Compose:
```bash
make docker-up-postgres
```

3. The application will automatically create the tables on startup.

## Running the Application

```bash
go run main.go
```

The service will start on `http://localhost:8091` (or the port specified in your configuration).

## API Endpoints

### Public Endpoints

#### Register User
```bash
POST /register
# or
POST /auth/register
Content-Type: application/json

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

#### Login
```bash
POST /login
# or
POST /auth/login
Content-Type: application/json

{
  "username": "john_doe",
  "password": "securepassword123"
}
```

### Google OAuth Authentication

#### Get Google OAuth URL
```bash
GET /auth/google/url
```

#### Google OAuth Login
```bash
POST /auth/google
Content-Type: application/json

{
  "code": "<google-oauth-code>"
}
```

### Protected Endpoints (Require JWT Token)

#### Get Current User Info
```bash
GET /me
# or
GET /api/v1/me
Authorization: Bearer <your-jwt-token>
```

#### Role Management
```bash
# Create role
POST /api/v1/roles
Authorization: Bearer <your-jwt-token>
Content-Type: application/json
{
  "role": "admin",
  "description": "Administrator role"
}

# Get all roles
GET /api/v1/roles
Authorization: Bearer <your-jwt-token>

# Get role by ID
GET /api/v1/roles/{id}
Authorization: Bearer <your-jwt-token>

# Update role
PUT /api/v1/roles/{id}
Authorization: Bearer <your-jwt-token>
Content-Type: application/json
{
  "role": "super_admin",
  "description": "Super administrator role"
}

# Delete role
DELETE /api/v1/roles/{id}
Authorization: Bearer <your-jwt-token>
```

### System Endpoints

#### Health Check
```bash
GET /health
```

#### Swagger Documentation
Visit `http://localhost:8091/swagger/index.html` to view the interactive API documentation.

## Example Usage

### 1. Register a new user:
```bash
curl -X POST http://localhost:8091/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "email": "john.doe@example.com",
    "mobile": "+1234567890",
    "display_name": "John Doe",
    "password": "securepassword123",
    "country": "US",
    "role": "user"
  }'
```

### 2. Login to get a JWT token:
```bash
curl -X POST http://localhost:8091/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "password": "securepassword123"
  }'
```

Response:
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

### 3. Access protected endpoint:
```bash
curl -X GET http://localhost:8091/me \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

## JWT Token Structure

The JWT tokens contain the following claims:

- `sub`: User ID (subject)
- `username`: Username
- `exp`: Expiration time (default: 15 minutes)
- `iss`: Issuer ("iam-service")
- `iat`: Issued at timestamp
- `nbf`: Not before timestamp

## User Model

The `users` table contains:

- `id`: UUID primary key
- `username`: Unique username
- `email`: Unique email address
- `mobile`: Unique mobile number
- `display_name`: Display name
- `password`: Hashed password (bcrypt)
- `country`: ISO country code
- `role`: User role (default: "user") - **Hidden from API responses**
- `created_at`: Creation timestamp
- `updated_at`: Last update timestamp

**Note**: The `role` field is stored in the database but is not returned in any API responses for security reasons.

## Development

### Adding Custom Claims

You can add custom claims to JWT tokens by modifying the login handler:

```go
customClaims := map[string]interface{}{
    "department": "engineering",
    "permissions": []string{"read", "write"},
}

token, err := jwtManager.GenerateTokenWithCustomClaims(
    user.ID.String(),
    user.Username,
    user.Role,
    customClaims,
)
```

### Extending the API

1. Add new handlers in `internal/auth/handler.go`
2. Add Swagger annotations to your handlers
3. Update routes in `main.go`
4. Regenerate Swagger docs: `swag init`

## Error Handling

The API returns structured error responses:

```json
{
  "error": "error_code",
  "message": "Human readable error message"
}
```

Common error codes:
- `validation_error`: Request validation failed
- `authentication_failed`: Invalid credentials
- `creation_failed`: User creation failed
- `invalid_token`: Invalid or expired JWT token
- `user_not_found`: User not found

## Security Considerations

- Passwords are hashed using bcrypt with cost factor 12
- JWT tokens are signed with HMAC-SHA256
- Default token expiration is 15 minutes
- CORS is enabled for all origins (configure for production)
- Database connections use SSL (configurable)

## Production Deployment

1. Set strong JWT secret in environment variables
2. Configure proper CORS settings
3. Use SSL/TLS for database connections
4. Set up proper logging and monitoring
5. Configure reverse proxy (nginx/Apache)
6. Use environment-specific configuration

## License

This project is licensed under the MIT License.
