package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret        string
	DatabaseURL      string
	DatabaseType     string
	TokenExpMinutes  int
	Port             string
	TokenExpDuration time.Duration

	// Google OAuth Configuration
	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURL  string
}

func Load() *Config {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	config := &Config{
		JWTSecret:       getEnv("JWT_SECRET", "default-secret-key"),
		DatabaseURL:     getEnv("DATABASE_URL", "fitflow_iam_user:password@tcp(localhost:3306)/fitflow_iam_db?charset=utf8mb4&parseTime=True&loc=Local"),
		DatabaseType:    getEnv("DATABASE_TYPE", "mysql"),
		TokenExpMinutes: getEnvAsInt("TOKEN_EXP_MINUTES", 15),
		Port:            getEnv("PORT", "8091"),

		// Google OAuth Configuration
		GoogleClientID:     getEnv("GOOGLE_CLIENT_ID", ""),
		GoogleClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
		GoogleRedirectURL:  getEnv("GOOGLE_REDIRECT_URL", "http://localhost:3000/auth/google/callback"),
	}

	config.TokenExpDuration = time.Duration(config.TokenExpMinutes) * time.Minute

	return config
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
