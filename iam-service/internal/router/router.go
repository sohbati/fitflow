package router

import (
	"fmt"
	"iam-service/internal/auth"
	"iam-service/internal/middleware"
	"iam-service/internal/role"
	"iam-service/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	authHandler *auth.Handler
	roleHandler *role.Handler
	jwtManager  *jwt.JWTManager
}

type RouterConfig struct {
	EnableCORS    bool
	EnableSwagger bool
	EnableHealth  bool
	CORSOrigins   []string
	CORSMethods   []string
	CORSHeaders   []string
}

func NewRouter(authHandler *auth.Handler, roleHandler *role.Handler, jwtManager *jwt.JWTManager) *Router {
	return &Router{
		authHandler: authHandler,
		roleHandler: roleHandler,
		jwtManager:  jwtManager,
	}
}

func (r *Router) SetupRoutes() *gin.Engine {
	return r.SetupRoutesWithConfig(&RouterConfig{
		EnableCORS:    true,
		EnableSwagger: true,
		EnableHealth:  true,
		CORSOrigins:   []string{"*"},
		CORSMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		CORSHeaders:   []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", "sec-ch-ua", "sec-ch-ua-mobile", "sec-ch-ua-platform", "User-Agent", "Referer"},
	})
}

func (r *Router) SetupRoutesWithConfig(config *RouterConfig) *gin.Engine {
	router := gin.Default()

	// Add CORS middleware if enabled
	if config.EnableCORS {
		r.setupCORS(router, config)
	}

	// Health check endpoint if enabled
	if config.EnableHealth {
		r.setupHealthCheck(router)
	}

	// Swagger documentation if enabled
	if config.EnableSwagger {
		r.setupSwagger(router)
	}

	// Public routes
	r.setupPublicRoutes(router)

	// Protected routes
	r.setupProtectedRoutes(router)

	return router
}

func (r *Router) setupCORS(router *gin.Engine, config *RouterConfig) {
	router.Use(func(c *gin.Context) {
		// Set CORS headers
		if len(config.CORSOrigins) == 1 && config.CORSOrigins[0] == "*" {
			c.Header("Access-Control-Allow-Origin", "*")
		} else {
			origin := c.Request.Header.Get("Origin")
			for _, allowedOrigin := range config.CORSOrigins {
				if origin == allowedOrigin {
					c.Header("Access-Control-Allow-Origin", origin)
					break
				}
			}
		}

		c.Header("Access-Control-Allow-Methods", strings.Join(config.CORSMethods, ", "))
		c.Header("Access-Control-Allow-Headers", strings.Join(config.CORSHeaders, ", "))

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
}

func (r *Router) setupHealthCheck(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "iam-service",
		})
	})
}

func (r *Router) setupSwagger(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (r *Router) setupPublicRoutes(router *gin.Engine) {
	// Authentication routes
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", r.authHandler.Register)
		authGroup.POST("/login", r.authHandler.Login)

		// Google OAuth routes
		if r.authHandler.GoogleAuthHandler != nil {
			fmt.Printf("Router: Google OAuth handler is available\n")
			authGroup.POST("/google", func(c *gin.Context) {
				fmt.Printf("Router: POST /auth/google called\n")
				r.authHandler.GoogleAuthHandler.GoogleLogin(c)
			})
			authGroup.GET("/google/url", r.authHandler.GoogleAuthHandler.GetGoogleAuthURL)
		} else {
			fmt.Printf("Router: Google OAuth handler is nil\n")
		}
	}

	// Legacy routes for backward compatibility
	router.POST("/register", r.authHandler.Register)
	router.POST("/login", r.authHandler.Login)
}

func (r *Router) setupProtectedRoutes(router *gin.Engine) {
	// API v1 routes
	v1 := router.Group("/api/v1")
	v1.Use(middleware.JWTAuth(r.jwtManager))
	{
		v1.GET("/me", r.authHandler.GetMe)

		// Role management routes
		roles := v1.Group("/roles")
		{
			roles.POST("", r.roleHandler.CreateRole)
			roles.GET("", r.roleHandler.GetAllRoles)
			roles.GET("/:id", r.roleHandler.GetRole)
			roles.PUT("/:id", r.roleHandler.UpdateRole)
			roles.DELETE("/:id", r.roleHandler.DeleteRole)
		}
	}

	// Legacy protected routes for backward compatibility
	protected := router.Group("/")
	protected.Use(middleware.JWTAuth(r.jwtManager))
	{
		protected.GET("/me", r.authHandler.GetMe)
	}
}
