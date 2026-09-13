package config

import (
	"fmt"

	"GRPC/knight/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=mydb port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database connection failed")
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic(fmt.Sprintf("Failed to migrate database: %v", err))
	}

	DB = db
}
