package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"server-v2/models"
)

var DB *gorm.DB

func InitDatabase(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print("Migrating database...")
	migrateError := db.AutoMigrate(&models.Role{}, &models.User{}, &models.Oil{}, &models.VehicleType{}, &models.Vehicle{}, &models.Transaction{})
	if migrateError != nil {
		log.Fatalln(err)
		return nil
	}

	DB = db

	return nil

}
