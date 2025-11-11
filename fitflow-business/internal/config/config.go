package config

import (
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL  string
	DatabaseType string
	Port         string
}

func Load() *Config {
	godotenv.Load()
	
	return &Config{
		DatabaseURL:  getEnv("DATABASE_URL", "postgres://fitflow_business_user:password@localhost:5432/fitflow_business_db?sslmode=disable"),
		DatabaseType: getEnv("DATABASE_TYPE", "postgres"),
		Port:         getEnv("PORT", "8092"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

