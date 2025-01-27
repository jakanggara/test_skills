package main

import (
	"log"
	"skill_test/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Automatically migrate your schema
	db.AutoMigrate(&models.DbSource{}, &models.DbDestination{})
}
