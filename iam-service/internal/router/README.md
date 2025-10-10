# Router Package

The router package provides a clean and organized way to manage HTTP routes for the IAM service. It separates routing logic from the main application and provides configuration options for different environments.

## Features

- **Modular Route Organization** - Routes are grouped by functionality
- **Configurable Middleware** - CORS, health checks, and Swagger can be enabled/disabled
- **API Versioning** - Support for versioned API endpoints
- **Backward Compatibility** - Legacy routes are maintained
- **Extensible Design** - Easy to add new route groups and middleware

## Usage

### Basic Usage

```go
// Create router with default configuration
appRouter := router.NewRouter(authHandler, jwtManager)
ginRouter := appRouter.SetupRoutes()
```

### Advanced Configuration

```go
// Create router with custom configuration
config := &router.RouterConfig{
    EnableCORS:    true,
    EnableSwagger: true,
    EnableHealth:  true,
    CORSOrigins:   []string{"http://localhost:3000", "https://myapp.com"},
    CORSMethods:   []string{"GET", "POST", "PUT", "DELETE"},
    CORSHeaders:   []string{"Content-Type", "Authorization"},
}

appRouter := router.NewRouter(authHandler, jwtManager)
ginRouter := appRouter.SetupRoutesWithConfig(config)
```

## Route Structure

### Public Routes

- **Legacy Routes**:
  - `POST /register` - User registration
  - `POST /login` - User authentication

- **Grouped Routes**:
  - `POST /auth/register` - User registration
  - `POST /auth/login` - User authentication

### Protected Routes

- **Legacy Routes**:
  - `GET /me` - Get current user info

- **Versioned API Routes**:
  - `GET /api/v1/me` - Get current user info

### System Routes

- `GET /health` - Health check endpoint
- `GET /swagger/*any` - Swagger documentation

## Configuration Options

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `EnableCORS` | bool | true | Enable CORS middleware |
| `EnableSwagger` | bool | true | Enable Swagger documentation |
| `EnableHealth` | bool | true | Enable health check endpoint |
| `CORSOrigins` | []string | ["*"] | Allowed CORS origins |
| `CORSMethods` | []string | ["GET", "POST", "PUT", "DELETE", "OPTIONS"] | Allowed HTTP methods |
| `CORSHeaders` | []string | ["Origin", "Content-Type", "Accept", "Authorization"] | Allowed headers |

## Adding New Routes

### 1. Add Route Group

```go
func (r *Router) setupPublicRoutes(router *gin.Engine) {
    // Existing routes...
    
    // Add new route group
    publicGroup := router.Group("/public")
    {
        publicGroup.GET("/info", r.infoHandler.GetInfo)
    }
}
```

### 2. Add Protected Routes

```go
func (r *Router) setupProtectedRoutes(router *gin.Engine) {
    // Existing protected routes...
    
    // Add new protected route group
    adminGroup := router.Group("/admin")
    adminGroup.Use(middleware.JWTAuth(r.jwtManager))
    adminGroup.Use(middleware.AdminOnly()) // Custom middleware
    {
        adminGroup.GET("/users", r.adminHandler.GetAllUsers)
    }
}
```

### 3. Add Custom Middleware

```go
func (r *Router) setupCustomMiddleware(router *gin.Engine) {
    router.Use(gin.Logger())
    router.Use(gin.Recovery())
    router.Use(middleware.RequestID())
}
```

## Environment-Specific Configuration

### Development

```go
devConfig := &router.RouterConfig{
    EnableCORS:    true,
    EnableSwagger: true,
    EnableHealth:  true,
    CORSOrigins:   []string{"*"},
}
```

### Production

```go
prodConfig := &router.RouterConfig{
    EnableCORS:    true,
    EnableSwagger: false, // Disable in production
    EnableHealth:  true,
    CORSOrigins:   []string{"https://myapp.com"},
}
```

## Testing

The router can be easily tested by creating a test instance:

```go
func TestRouter(t *testing.T) {
    // Setup test dependencies
    authHandler := &mockAuthHandler{}
    jwtManager := &mockJWTManager{}
    
    // Create router
    appRouter := router.NewRouter(authHandler, jwtManager)
    ginRouter := appRouter.SetupRoutes()
    
    // Test routes
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/health", nil)
    ginRouter.ServeHTTP(w, req)
    
    assert.Equal(t, 200, w.Code)
}
```

## Benefits

1. **Separation of Concerns** - Routing logic is isolated from business logic
2. **Maintainability** - Easy to modify routes without touching main.go
3. **Testability** - Router can be tested independently
4. **Flexibility** - Configuration allows for different environments
5. **Scalability** - Easy to add new route groups and middleware
6. **Documentation** - Clear structure makes the API self-documenting
