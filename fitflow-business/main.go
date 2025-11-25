package main

import (
	"log"

	"fitflow-business/internal/config"
	"fitflow-business/internal/db"
	"fitflow-business/internal/router"

	"gorm.io/gorm"
)

// @title FitFlow Business Service API
// @version 1.0
// @description Business Logic Service for FitFlow Application
// @host localhost:8092
// @BasePath /

type Application struct {
	config   *config.Config
	database *gorm.DB
	router   *router.Router
}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) Initialize() error {
	app.config = config.Load()
	log.Println("Configuration loaded")

	log.Println("Connecting to database...")
	database, err := db.Connect(app.config.DatabaseURL, app.config.DatabaseType)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}
	app.database = database
	log.Println("Database connection established and verified")

	app.router = router.NewRouter(app.database)
	log.Println("Router configured")

	return nil
}

func (app *Application) Run() error {
	ginRouter := app.router.SetupRoutes()
	log.Printf("Starting FitFlow Business Service on port %s", app.config.Port)
	return ginRouter.Run(":" + app.config.Port)
}

func main() {
	app := NewApplication()
	if err := app.Initialize(); err != nil {
		log.Fatalf("Failed to initialize: %v", err)
	}
	if err := app.Run(); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
