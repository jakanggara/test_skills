package main

import (
	"log"
	"skill_test/api"
	"skill_test/data"
	"skill_test/models"

	"github.com/gin-gonic/gin"
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

	totalData := 500
	data.Seed(db, uint(totalData))

	r := gin.Default()
	api.InitRoutes(r, db)

	r.Run(":8001")
}
