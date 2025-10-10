# Normalized Authentication System

## Overview

The IAM service has been refactored to support multiple authentication providers using a normalized database schema. This allows for easy addition of new authentication methods like Facebook, Apple, mobile OTP, and email OTP.

## Database Schema

### 1. Users Table (Normalized)
```sql
CREATE TABLE users (
    id CHAR(36) PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    display_name VARCHAR(255) NOT NULL,
    avatar_url VARCHAR(500),
    country VARCHAR(255),
    role VARCHAR(255) DEFAULT 'user',
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

### 2. Auth Providers Table
```sql
CREATE TABLE auth_providers (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    display_name VARCHAR(100) NOT NULL,
    description TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

**Default Providers:**
- `local` - Username/password authentication
- `google` - Google OAuth 2.0
- `facebook` - Facebook OAuth 2.0 (ready for implementation)
- `apple` - Apple Sign In (ready for implementation)
- `mobile_otp` - Mobile phone OTP (ready for implementation)
- `email_otp` - Email OTP (ready for implementation)

### 3. User Auth Table (Links users to providers)
```sql
CREATE TABLE user_auth (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    auth_provider_id CHAR(36) NOT NULL,
    
    -- Provider-specific identifiers
    provider_user_id VARCHAR(255), -- Google ID, Facebook ID, etc.
    provider_email VARCHAR(255),   -- Email from provider
    username VARCHAR(255),         -- For local authentication
    
    -- Provider-specific data (JSON)
    provider_data JSON,           -- Store additional provider data
    
    -- Verification status
    is_verified BOOLEAN DEFAULT FALSE,
    verified_at TIMESTAMP NULL,
    
    -- OTP fields (for mobile/email OTP)
    otp_code VARCHAR(10) NULL,
    otp_expires_at TIMESTAMP NULL,
    otp_attempts INT DEFAULT 0,
    
    -- Timestamps
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    -- Foreign keys
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (auth_provider_id) REFERENCES auth_providers(id) ON DELETE CASCADE,
    
    -- Unique constraints
    UNIQUE KEY unique_user_provider (user_id, auth_provider_id),
    UNIQUE KEY unique_provider_user_id (auth_provider_id, provider_user_id)
);
```

## Architecture

### 1. Models
- **User**: Core user information (email, display name, avatar, etc.)
- **AuthProvider**: Available authentication methods
- **UserAuth**: Links users to their authentication providers

### 2. Services
- **AuthProviderService**: Manages authentication providers
- **UserAuthService**: Manages user-provider relationships
- **LocalAuthHandler**: Handles username/password authentication
- **GoogleAuthHandler**: Handles Google OAuth authentication

### 3. Repositories
- **AuthProviderRepository**: Database operations for auth providers
- **UserAuthRepository**: Database operations for user auth relationships

## API Endpoints

### Authentication
- `POST /auth/register` - Register with username/password
- `POST /auth/login` - Login with username/password
- `POST /auth/google` - Login with Google OAuth
- `GET /auth/google/url` - Get Google OAuth URL

### User Management
- `GET /me` - Get current user info
- `GET /api/v1/me` - Get current user info (v1)

### Role Management
- `POST /api/v1/roles` - Create role
- `GET /api/v1/roles` - Get all roles
- `GET /api/v1/roles/{id}` - Get role by ID
- `PUT /api/v1/roles/{id}` - Update role
- `DELETE /api/v1/roles/{id}` - Delete role

## Adding New Authentication Providers

### 1. Database Setup
Add the new provider to the `auth_providers` table:
```sql
INSERT INTO auth_providers (name, display_name, description) VALUES
('facebook', 'Facebook OAuth', 'Facebook OAuth 2.0 authentication');
```

### 2. Create Handler
Create a new handler similar to `GoogleAuthHandler`:
```go
type FacebookAuthHandler struct {
    userService       user.Service
    userAuthService   UserAuthService
    authProviderService AuthProviderService
    jwtManager        *jwt.JWTManager
    config            *oauth2.Config
}
```

### 3. Add Routes
Add routes to the router:
```go
authGroup.POST("/facebook", r.authHandler.FacebookAuthHandler.FacebookLogin)
authGroup.GET("/facebook/url", r.authHandler.FacebookAuthHandler.GetFacebookAuthURL)
```

### 4. Update Configuration
Add Facebook OAuth configuration to environment variables:
```env
FACEBOOK_CLIENT_ID=your_facebook_client_id
FACEBOOK_CLIENT_SECRET=your_facebook_client_secret
FACEBOOK_REDIRECT_URL=http://localhost:8091/auth/facebook/callback
```

## OTP Authentication (Mobile/Email)

The system includes built-in support for OTP authentication:

### Generate OTP
```go
userAuth, err := userAuthService.GenerateOTP(userID, providerID, email)
```

### Verify OTP
```go
userAuth, err := userAuthService.VerifyOTP(userID, providerID, otpCode)
```

### OTP Features
- 6-digit OTP codes
- 10-minute expiration
- 3-attempt limit
- Automatic cleanup after verification

## Database Setup (Early Development Phase)

Since this project is in early development phase, we use direct schema changes instead of migrations:

### 1. Database Initialization
The database schema is managed using manual SQL migrations (GORM AutoMigrate is disabled).

### 2. Schema Files (Reference Only)
The schema files in the `database/` folder are for reference:
- `1_users.sql` - Normalized users table
- `2_roles.sql` - Roles table
- `3_auth_providers.sql` - Auth providers table
- `4_user_auth.sql` - User auth relationships table

### 3. Manual SQL Migrations
The application uses manual SQL migrations (AutoMigrate disabled):
```go
func Migrate(db *gorm.DB) error {
    // AutoMigrate disabled for both development and production
    // Use manual SQL migrations instead
    log.Println("AutoMigrate disabled - using manual SQL migrations")
    return nil
}
```

## Benefits

1. **Scalability**: Easy to add new authentication providers
2. **Flexibility**: Users can have multiple authentication methods
3. **Security**: Provider-specific data is isolated
4. **Maintainability**: Clear separation of concerns
5. **Extensibility**: OTP support built-in for mobile/email authentication

## Security Considerations

1. **Password Storage**: Passwords are hashed using bcrypt
2. **JWT Tokens**: Secure token generation with configurable expiration
3. **OTP Security**: Rate limiting and expiration for OTP codes
4. **Provider Data**: Sensitive provider data stored in JSON field
5. **Foreign Keys**: Cascading deletes maintain data integrity

## Future Enhancements

1. **Facebook OAuth**: Ready for implementation
2. **Apple Sign In**: Ready for implementation
3. **Mobile OTP**: SMS integration needed
4. **Email OTP**: Email service integration needed
5. **Two-Factor Authentication**: Combine multiple providers
6. **Social Login Linking**: Link multiple social accounts to one user
7. **Account Recovery**: Password reset via OTP
8. **Audit Logging**: Track authentication events
9. **Rate Limiting**: Prevent brute force attacks
10. **Device Management**: Track and manage user devices
