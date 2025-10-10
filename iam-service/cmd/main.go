package main

import (
	"log"

	"iam-service/internal/auth"
	"iam-service/internal/config"
	"iam-service/internal/db"
	"iam-service/internal/role"
	"iam-service/internal/router"
	"iam-service/internal/user"
	"iam-service/pkg/jwt"

	"gorm.io/gorm"
)

// @title IAM Service API
// @version 1.0
// @description Identity and Access Management Service
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8091
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// Application represents the main application structure
type Application struct {
	config      *config.Config
	database    *gorm.DB
	jwtManager  *jwt.JWTManager
	userRepo    user.Repository
	userService user.Service
	roleRepo    role.Repository
	roleService role.Service
	authHandler *auth.Handler
	roleHandler *role.Handler
	router      *router.Router
}

// NewApplication creates a new application instance
func NewApplication() *Application {
	return &Application{}
}

// Initialize sets up all application dependencies
func (app *Application) Initialize() error {
	if err := app.loadConfiguration(); err != nil {
		return err
	}

	if err := app.setupDatabase(); err != nil {
		return err
	}

	if err := app.setupJWTManager(); err != nil {
		return err
	}

	if err := app.setupServices(); err != nil {
		return err
	}

	if err := app.setupHandlers(); err != nil {
		return err
	}

	if err := app.setupRouter(); err != nil {
		return err
	}

	return nil
}

// loadConfiguration loads application configuration
func (app *Application) loadConfiguration() error {
	app.config = config.Load()
	log.Println("Configuration loaded successfully")
	return nil
}

// setupDatabase initializes database connection and runs migrations
func (app *Application) setupDatabase() error {
	database, err := db.Connect(app.config.DatabaseURL, app.config.DatabaseType)
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

// setupJWTManager initializes JWT token manager
func (app *Application) setupJWTManager() error {
	app.jwtManager = jwt.NewJWTManager(app.config.JWTSecret, app.config.TokenExpDuration)
	log.Println("JWT manager initialized successfully")
	return nil
}

// setupServices initializes business logic services
func (app *Application) setupServices() error {
	app.userRepo = user.NewRepository(app.database)
	app.userService = user.NewService(app.userRepo)

	app.roleRepo = role.NewRepository(app.database)
	app.roleService = role.NewService(app.roleRepo)

	log.Println("Services initialized successfully")
	return nil
}

// setupHandlers initializes HTTP request handlers
func (app *Application) setupHandlers() error {
	app.authHandler = auth.NewHandler(app.userService, app.jwtManager)
	app.roleHandler = role.NewHandler(app.roleService)
	log.Println("Handlers initialized successfully")
	return nil
}

// setupRouter configures HTTP routing
func (app *Application) setupRouter() error {
	app.router = router.NewRouter(app.authHandler, app.roleHandler, app.jwtManager)
	log.Println("Router configured successfully")
	return nil
}

// Run starts the HTTP server
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

// Shutdown gracefully shuts down the application
func (app *Application) Shutdown() {
	if app.database != nil {
		log.Println("Shutting down database connection...")
		// Add database cleanup logic here if needed
	}
	log.Println("Application shutdown complete")
}

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
