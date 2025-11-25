package db

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func Connect(databaseURL, databaseType string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	switch databaseType {
	case "mysql":
		db, err = gorm.Open(mysql.Open(databaseURL), &gorm.Config{})
	case "postgres":
		db, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	default:
		log.Fatalf("Unsupported database type: %s", databaseType)
	}

	if err != nil {
		return nil, err
	}

	log.Printf("Database connected successfully using %s", databaseType)
	
	// Verify connection by pinging the database
	if err := CheckConnection(db); err != nil {
		return nil, err
	}

	return db, nil
}

// CheckConnection verifies that the database connection is working by pinging it
func CheckConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Ping the database to verify connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		return err
	}

	log.Println("Database connection verified successfully")
	return nil
}

func Migrate(db *gorm.DB) error {
	log.Println("Database migrations disabled - using manual SQL migrations")
	return nil
}
