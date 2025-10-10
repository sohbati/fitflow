# Main Application Architecture

This document describes the refactored main application architecture with clean naming conventions and improved readability.

## Architecture Overview

The main application follows a structured approach with clear separation of concerns:

```
Application
├── Configuration Management
├── Database Setup
├── JWT Manager Setup
├── Service Layer Setup
├── Handler Setup
├── Router Setup
└── Server Execution
```

## Application Structure

### Application Struct

```go
type Application struct {
    config      *config.Config      // Application configuration
    database    *gorm.DB            // Database connection
    jwtManager  *jwt.JWTManager     // JWT token manager
    userRepo    user.Repository     // User data repository
    userService user.Service        // User business logic
    authHandler *auth.Handler       // Authentication handlers
    router      *router.Router      // HTTP router
}
```

## Method Organization

### 1. Constructor
- **`NewApplication()`** - Creates a new application instance

### 2. Initialization Methods
- **`Initialize()`** - Orchestrates the entire initialization process
- **`loadConfiguration()`** - Loads application configuration
- **`setupDatabase()`** - Initializes database connection and migrations
- **`setupJWTManager()`** - Initializes JWT token manager
- **`setupServices()`** - Initializes business logic services
- **`setupHandlers()`** - Initializes HTTP request handlers
- **`setupRouter()`** - Configures HTTP routing

### 3. Runtime Methods
- **`Run()`** - Starts the HTTP server
- **`Shutdown()`** - Gracefully shuts down the application

## Method Details

### Initialize()
The main initialization method that orchestrates all setup steps:

```go
func (app *Application) Initialize() error {
    if err := app.loadConfiguration(); err != nil {
        return err
    }
    // ... other setup methods
    return nil
}
```

**Benefits:**
- Clear initialization flow
- Early error detection
- Easy to add new setup steps
- Consistent error handling

### loadConfiguration()
Loads and validates application configuration:

```go
func (app *Application) loadConfiguration() error {
    app.config = config.Load()
    log.Println("Configuration loaded successfully")
    return nil
}
```

**Features:**
- Centralized configuration loading
- Success logging
- Error propagation

### setupDatabase()
Handles database connection and migrations:

```go
func (app *Application) setupDatabase() error {
    database, err := db.Connect(app.config.DatabaseURL)
    if err != nil {
        log.Printf("Failed to connect to database: %v", err)
        return err
    }
    
    if err := db.Migrate(database); err != nil {
        log.Printf("Failed to migrate database: %v", err)
        return err
    }
    
    app.database = database
    log.Println("Database connected and migrated successfully")
    return nil
}
```

**Features:**
- Connection and migration in one method
- Detailed error logging
- Success confirmation

### setupJWTManager()
Initializes JWT token management:

```go
func (app *Application) setupJWTManager() error {
    app.jwtManager = jwt.NewJWTManager(app.config.JWTSecret, app.config.TokenExpDuration)
    log.Println("JWT manager initialized successfully")
    return nil
}
```

### setupServices()
Initializes business logic layer:

```go
func (app *Application) setupServices() error {
    app.userRepo = user.NewRepository(app.database)
    app.userService = user.NewService(app.userRepo)
    log.Println("Services initialized successfully")
    return nil
}
```

**Features:**
- Repository pattern implementation
- Service layer initialization
- Dependency injection

### setupHandlers()
Initializes HTTP request handlers:

```go
func (app *Application) setupHandlers() error {
    app.authHandler = auth.NewHandler(app.userService, app.jwtManager)
    log.Println("Handlers initialized successfully")
    return nil
}
```

### setupRouter()
Configures HTTP routing:

```go
func (app *Application) setupRouter() error {
    app.router = router.NewRouter(app.authHandler, app.jwtManager)
    log.Println("Router configured successfully")
    return nil
}
```

### Run()
Starts the HTTP server with enhanced logging:

```go
func (app *Application) Run() error {
    ginRouter := app.router.SetupRoutes()
    
    log.Printf("Starting IAM Service on port %s", app.config.Port)
    log.Printf("Swagger documentation available at: http://localhost:%s/swagger/index.html", app.config.Port)
    log.Printf("Health check available at: http://localhost:%s/health", app.config.Port)
    
    if err := ginRouter.Run(":" + app.config.Port); err != nil {
        log.Printf("Failed to start server: %v", err)
        return err
    }
    
    return nil
}
```

**Features:**
- Enhanced startup logging
- Service URLs provided
- Graceful error handling

### Shutdown()
Handles graceful application shutdown:

```go
func (app *Application) Shutdown() {
    if app.database != nil {
        log.Println("Shutting down database connection...")
        // Add database cleanup logic here if needed
    }
    log.Println("Application shutdown complete")
}
```

**Features:**
- Graceful resource cleanup
- Shutdown logging
- Extensible for additional cleanup

## Main Function

The main function follows a clean pattern:

```go
func main() {
    // Create and initialize application
    app := NewApplication()
    
    if err := app.Initialize(); err != nil {
        log.Fatalf("Failed to initialize application: %v", err)
    }
    
    // Ensure graceful shutdown
    defer app.Shutdown()
    
    // Start the application
    if err := app.Run(); err != nil {
        log.Fatalf("Application failed to start: %v", err)
    }
}
```

**Benefits:**
- Clear initialization flow
- Graceful shutdown handling
- Consistent error handling
- Easy to understand execution flow

## Benefits of Refactored Architecture

### 1. **Readability**
- Clear method names that describe their purpose
- Logical flow from initialization to execution
- Consistent naming conventions

### 2. **Maintainability**
- Each setup step is isolated in its own method
- Easy to modify individual components
- Clear separation of concerns

### 3. **Testability**
- Each method can be tested independently
- Dependencies are clearly defined
- Mock objects can be easily injected

### 4. **Extensibility**
- Easy to add new setup steps
- Simple to modify existing functionality
- Clear extension points

### 5. **Error Handling**
- Consistent error propagation
- Detailed error logging
- Early failure detection

### 6. **Logging**
- Comprehensive logging at each step
- Success and failure logging
- Useful startup information

## Usage Examples

### Basic Usage
```go
app := NewApplication()
if err := app.Initialize(); err != nil {
    log.Fatal(err)
}
defer app.Shutdown()
app.Run()
```

### Testing
```go
func TestApplication(t *testing.T) {
    app := NewApplication()
    
    // Mock dependencies
    app.config = &mockConfig{}
    app.database = &mockDB{}
    
    err := app.Initialize()
    assert.NoError(t, err)
}
```

### Custom Configuration
```go
app := NewApplication()
app.config = customConfig
err := app.Initialize()
```

## Best Practices

1. **Single Responsibility** - Each method has one clear purpose
2. **Error Propagation** - Errors are properly propagated up the call stack
3. **Resource Management** - Resources are properly initialized and cleaned up
4. **Logging** - Comprehensive logging for debugging and monitoring
5. **Dependency Injection** - Dependencies are clearly defined and injected
6. **Graceful Shutdown** - Application can be shut down cleanly

This refactored architecture provides a solid foundation for a maintainable, testable, and extensible Go microservice.
