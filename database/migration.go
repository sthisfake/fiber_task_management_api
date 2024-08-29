package database

import (
	"log"
	"task/models"
)

func Migrate() {
	if err := DB.AutoMigrate(&models.User{}, &models.Task{}); err != nil {
		log.Fatalf("could not migrate the database: %v", err)
	}
}
