package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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
	return db, nil
}

func Migrate(db *gorm.DB) error {
	log.Println("Database migrations disabled - using manual SQL migrations")
	return nil
}
