# Auth Package Architecture

This document describes the refactored authentication package structure with improved modularity and maintainability.

## Package Structure

```
internal/auth/
├── handler.go      # Main handler struct and constructor
├── register.go     # User registration logic
├── login.go        # User authentication logic
├── profile.go      # User profile management
├── types.go        # Common types and DTOs
└── README.md       # This documentation
```

## Architecture Overview

The auth package follows a modular approach where each file handles a specific authentication concern:

### 1. **handler.go** - Main Handler
- Defines the main `Handler` struct
- Contains constructor `NewHandler()`
- Manages dependencies (userService, jwtManager)

### 2. **register.go** - User Registration
- `RegisterRequest` struct
- `Register()` handler method
- User creation logic and validation

### 3. **login.go** - User Authentication
- `LoginRequest` struct
- `AuthResponse` struct
- `Login()` handler method
- JWT token generation

### 4. **profile.go** - User Profile Management
- `GetMe()` handler method
- User profile retrieval logic

### 5. **types.go** - Common Types
- `ErrorResponse` struct
- `SuccessResponse` struct
- Shared response types

## Benefits of Refactoring

### 1. **Separation of Concerns**
- Each file handles a specific authentication feature
- Clear boundaries between different functionalities
- Easier to locate and modify specific features

### 2. **Maintainability**
- Smaller, focused files are easier to understand
- Changes to one feature don't affect others
- Reduced cognitive load when working on specific features

### 3. **Testability**
- Individual features can be tested in isolation
- Mock dependencies more easily
- Better test coverage and organization

### 4. **Scalability**
- Easy to add new authentication features
- Clear pattern for extending functionality
- Consistent structure across all handlers

### 5. **Code Organization**
- Logical grouping of related functionality
- Consistent naming conventions
- Better code navigation

## File Details

### handler.go
```go
type Handler struct {
    userService user.Service
    jwtManager  *jwt.JWTManager
}

func NewHandler(userService user.Service, jwtManager *jwt.JWTManager) *Handler
```

**Responsibilities:**
- Define main handler structure
- Manage dependencies
- Provide constructor for dependency injection

### register.go
```go
type RegisterRequest struct {
    Username    string `json:"username" binding:"required"`
    Email       string `json:"email" binding:"required,email"`
    Mobile      string `json:"mobile" binding:"required"`
    DisplayName string `json:"display_name" binding:"required"`
    Password    string `json:"password" binding:"required"`
    Country     string `json:"country" binding:"required"`
    Role        string `json:"role,omitempty"`
}

func (h *Handler) Register(c *gin.Context)
```

**Responsibilities:**
- Handle user registration requests
- Validate registration data
- Create new users
- Handle registration errors

### login.go
```go
type LoginRequest struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
    Token string     `json:"token"`
    User  *user.User `json:"user"`
}

func (h *Handler) Login(c *gin.Context)
```

**Responsibilities:**
- Handle user authentication
- Validate login credentials
- Generate JWT tokens
- Return authentication response

### profile.go
```go
func (h *Handler) GetMe(c *gin.Context)
```

**Responsibilities:**
- Handle user profile requests
- Extract user ID from JWT context
- Retrieve user information
- Return user profile data

### types.go
```go
type ErrorResponse struct {
    Error   string `json:"error"`
    Message string `json:"message"`
}

type SuccessResponse struct {
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}
```

**Responsibilities:**
- Define common response types
- Provide consistent error handling
- Standardize API responses

## Usage Examples

### Creating Handler
```go
authHandler := auth.NewHandler(userService, jwtManager)
```

### Registering Routes
```go
// Public routes
router.POST("/register", authHandler.Register)
router.POST("/login", authHandler.Login)

// Protected routes
protected := router.Group("/")
protected.Use(middleware.JWTAuth(jwtManager))
{
    protected.GET("/me", authHandler.GetMe)
}
```

### Testing Individual Features
```go
func TestRegister(t *testing.T) {
    // Test registration logic in isolation
    handler := auth.NewHandler(mockUserService, mockJWTManager)
    // ... test implementation
}

func TestLogin(t *testing.T) {
    // Test login logic in isolation
    handler := auth.NewHandler(mockUserService, mockJWTManager)
    // ... test implementation
}
```

## API Endpoints

### Public Endpoints
- **POST /register** - User registration
- **POST /login** - User authentication

### Protected Endpoints
- **GET /me** - Get current user profile

## Error Handling

All handlers use consistent error handling through `ErrorResponse`:

```go
c.JSON(http.StatusBadRequest, ErrorResponse{
    Error:   "validation_error",
    Message: err.Error(),
})
```

## Success Responses

All successful operations use `SuccessResponse`:

```go
c.JSON(http.StatusOK, SuccessResponse{
    Message: "Operation successful",
    Data:    result,
})
```

## Future Enhancements

### Easy to Add New Features
1. **Password Reset** - Add `password_reset.go`
2. **Email Verification** - Add `verification.go`
3. **Two-Factor Authentication** - Add `2fa.go`
4. **Social Login** - Add `social.go`
5. **Account Management** - Add `account.go`

### Example: Adding Password Reset
```go
// internal/auth/password_reset.go
type PasswordResetRequest struct {
    Email string `json:"email" binding:"required,email"`
}

func (h *Handler) RequestPasswordReset(c *gin.Context) {
    // Implementation
}

func (h *Handler) ResetPassword(c *gin.Context) {
    // Implementation
}
```

## Best Practices

### 1. **File Naming**
- Use descriptive names that indicate functionality
- Keep file names short but clear
- Use snake_case for file names

### 2. **Function Organization**
- One main handler function per file
- Related helper functions in the same file
- Keep functions focused and small

### 3. **Type Definitions**
- Define types close to where they're used
- Use common types in `types.go`
- Keep request/response types with their handlers

### 4. **Dependency Management**
- Inject dependencies through constructor
- Use interfaces for better testability
- Keep dependencies minimal

### 5. **Error Handling**
- Use consistent error response format
- Provide meaningful error messages
- Handle different error types appropriately

## Migration Guide

### From Monolithic Handler
1. **Extract Types** - Move common types to `types.go`
2. **Split Handlers** - Move each handler method to its own file
3. **Update Imports** - Ensure all files import necessary packages
4. **Test Changes** - Verify all functionality still works

### Adding New Features
1. **Create New File** - Add new handler file (e.g., `feature.go`)
2. **Define Types** - Add request/response types
3. **Implement Handler** - Add handler method
4. **Update Routes** - Register new routes
5. **Add Tests** - Create tests for new functionality

This refactored structure provides a solid foundation for a maintainable and scalable authentication system.
