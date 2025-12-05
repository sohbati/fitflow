package router

import (
	"fitflow-business/internal/handler"
	"fitflow-business/internal/middleware"
	"fitflow-business/internal/repository/impl"
	"fitflow-business/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// CORS middleware
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

type Router struct {
	database *gorm.DB
}

func NewRouter(database *gorm.DB) *Router {
	return &Router{database: database}
}

func (r *Router) SetupRoutes() *gin.Engine {
	router := gin.New()

	// Apply global error handler and recovery middleware
	router.Use(middleware.CustomRecovery())
	router.Use(middleware.GlobalErrorHandler())

	// Apply CORS middleware to all routes
	router.Use(corsMiddleware())

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "fitflow-business"})
	})

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Initialize repositories and services
	gymRepo := impl.NewGymRepository(r.database)
	gymLocationRepo := impl.NewGymLocationRepository(r.database)
	gymOwnerRepo := impl.NewGymOwnerRepository(r.database)
	trainerRepo := impl.NewTrainerRepository(r.database)
	gymTrainerRepo := impl.NewGymTrainerRepository(r.database)
	traineeRepo := impl.NewTraineeRepository(r.database)
	personRepo := impl.NewPersonRepository(r.database)
	profileRepo := impl.NewProfileRepository(r.database)

	gymService := service.NewGymService(gymRepo, gymLocationRepo, gymOwnerRepo, trainerRepo, gymTrainerRepo)
	gymHandler := handler.NewGymHandler(gymService)

	traineeService := service.NewTraineeService(traineeRepo)
	traineeHandler := handler.NewTraineeHandler(traineeService)

	personService := service.NewPersonService(personRepo)
	profileService := service.NewProfileService(profileRepo, personRepo, gymOwnerRepo, trainerRepo, traineeRepo)
	registrationService := service.NewRegistrationService(personService, gymService, gymOwnerRepo, trainerRepo, traineeRepo, profileService)
	personHandler := handler.NewPersonHandler(personService, registrationService)

	profileHandler := handler.NewProfileHandler(profileService)

	// Business API routes
	api := router.Group("/api/v1")
	{
		// Gym routes
		gyms := api.Group("/gyms")
		{
			gyms.GET("", gymHandler.GetGyms)
			gyms.GET("/search", gymHandler.SearchGyms)
			gyms.GET("/verified", gymHandler.GetVerifiedGyms)
			gyms.GET("/:id", gymHandler.GetGym)
			gyms.POST("", gymHandler.CreateGym)
			gyms.PUT("/:id", gymHandler.UpdateGym)
			gyms.DELETE("/:id", gymHandler.DeleteGym)

			// Gym locations
			gyms.GET("/:id/locations", gymHandler.GetGymLocations)
			gyms.POST("/:id/locations", gymHandler.CreateGymLocation)

			// Gym owners
			gyms.GET("/:id/owners", gymHandler.GetGymOwners)
			gyms.POST("/:id/owners", gymHandler.CreateGymOwner)
		}

		// Gym owner routes
		gymOwners := api.Group("/gym-owners")
		{
			gymOwners.GET("/user/:user_id", gymHandler.GetGymOwnerByUserID)
		}

		// Trainer routes
		trainers := api.Group("/trainers")
		{
			trainers.GET("", gymHandler.GetTrainers)
			trainers.GET("/user/:user_id", gymHandler.GetTrainerByUserID)
			trainers.GET("/:id", gymHandler.GetTrainer)
			trainers.POST("", gymHandler.CreateTrainer)
		}

		// Person routes
		persons := api.Group("/persons")
		{
			persons.GET("/check/:user_id", personHandler.CheckPersonExists)
			persons.POST("/register/gym-owner", personHandler.RegisterGymOwner)
			persons.POST("/register/trainer", personHandler.RegisterTrainer)
			persons.POST("/register/trainee", personHandler.RegisterTrainee)
		}

		// Trainee routes
		trainees := api.Group("/trainees")
		{
			trainees.GET("", traineeHandler.GetTrainees)
			trainees.GET("/search", traineeHandler.SearchTrainees)
			trainees.GET("/active", traineeHandler.GetActiveTrainees)
			trainees.GET("/expiring", traineeHandler.GetExpiringMemberships)
			trainees.GET("/expired", traineeHandler.GetExpiredMemberships)
			trainees.GET("/email/:email", traineeHandler.GetTraineesByEmail)
			trainees.GET("/phone/:phone", traineeHandler.GetTraineesByPhone)
			trainees.GET("/membership/:type", traineeHandler.GetTraineesByMembershipType)
			trainees.GET("/fitness/:level", traineeHandler.GetTraineesByFitnessLevel)
			trainees.GET("/user/:user_id", traineeHandler.GetTraineeByUserID)
			trainees.GET("/:id", traineeHandler.GetTrainee)
			trainees.GET("/:id/age", traineeHandler.CalculateAge)
			trainees.GET("/:id/bmi", traineeHandler.CalculateBMI)
			trainees.GET("/:id/membership/validate", traineeHandler.ValidateMembership)
			trainees.POST("", traineeHandler.CreateTrainee)
			trainees.PUT("/:id", traineeHandler.UpdateTrainee)
			trainees.PATCH("/:id/status", traineeHandler.UpdateMembershipStatus)
			trainees.DELETE("/:id", traineeHandler.DeleteTrainee)
		}

		// Profile routes
		profiles := api.Group("/profiles")
		{
			profiles.POST("", profileHandler.CreateProfile)
			profiles.POST("/sync/:user_id", profileHandler.SyncProfiles)
			profiles.GET("/type/:type", profileHandler.GetProfilesByType)
			profiles.GET("/:id", profileHandler.GetProfile)
			profiles.PUT("/:id", profileHandler.UpdateProfile)
			profiles.PATCH("/:id/set-default", profileHandler.SetDefaultProfile)
			profiles.PATCH("/:id/activate", profileHandler.ActivateProfile)
			profiles.PATCH("/:id/deactivate", profileHandler.DeactivateProfile)
			profiles.DELETE("/:id", profileHandler.DeleteProfile)

			// User-specific profile routes
			profiles.GET("/user/:user_id", profileHandler.GetProfilesByUserID)
			profiles.GET("/user/:user_id/default", profileHandler.GetDefaultProfile)
			profiles.GET("/user/:user_id/active", profileHandler.GetActiveProfiles)
			profiles.GET("/user/:user_id/type/:type", profileHandler.GetProfileByType)
		}

		// Programs
		programs := api.Group("/programs")
		{
			programs.GET("", r.getPrograms)
			programs.GET("/:id", r.getProgram)
			programs.POST("", r.createProgram)
			programs.PUT("/:id", r.updateProgram)
			programs.DELETE("/:id", r.deleteProgram)
		}

		// Exercises
		exercises := api.Group("/exercises")
		{
			exercises.GET("", r.getExercises)
			exercises.GET("/:id", r.getExercise)
			exercises.POST("", r.createExercise)
			exercises.PUT("/:id", r.updateExercise)
			exercises.DELETE("/:id", r.deleteExercise)
		}

		// Workouts
		workouts := api.Group("/workouts")
		{
			workouts.GET("", r.getWorkouts)
			workouts.GET("/:id", r.getWorkout)
			workouts.POST("", r.createWorkout)
			workouts.PUT("/:id", r.updateWorkout)
			workouts.DELETE("/:id", r.deleteWorkout)
		}

		// Sessions
		sessions := api.Group("/sessions")
		{
			sessions.GET("", r.getSessions)
			sessions.GET("/:id", r.getSession)
			sessions.POST("", r.createSession)
			sessions.PUT("/:id", r.updateSession)
			sessions.DELETE("/:id", r.deleteSession)
		}

		// Analytics
		analytics := api.Group("/analytics")
		{
			analytics.GET("/dashboard", r.getDashboardAnalytics)
			analytics.GET("/progress", r.getProgressAnalytics)
		}
	}

	return router
}

// Placeholder handlers
func (r *Router) getPrograms(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (r *Router) getProgram(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (r *Router) createProgram(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (r *Router) updateProgram(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (r *Router) deleteProgram(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (r *Router) getExercises(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (r *Router) getExercise(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (r *Router) createExercise(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (r *Router) updateExercise(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (r *Router) deleteExercise(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (r *Router) getWorkouts(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (r *Router) getWorkout(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (r *Router) createWorkout(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (r *Router) updateWorkout(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (r *Router) deleteWorkout(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (r *Router) getSessions(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (r *Router) getSession(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (r *Router) createSession(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (r *Router) updateSession(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (r *Router) deleteSession(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (r *Router) getDashboardAnalytics(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}

func (r *Router) getProgressAnalytics(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
}
