package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(databaseURL, databaseType string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	switch databaseType {
	case "mysql":
		db, err = gorm.Open(mysql.Open(databaseURL), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	case "postgres":
		db, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	default:
		log.Fatalf("Unsupported database type: %s. Supported types: mysql, postgres", databaseType)
	}

	if err != nil {
		return nil, err
	}

	log.Printf("Database connected successfully using %s", databaseType)
	return db, nil
}

func Migrate(db *gorm.DB) error {
	// AutoMigrate disabled for both development and production
	// Use manual SQL migrations instead
	log.Println("AutoMigrate disabled - using manual SQL migrations")
	return nil
}
