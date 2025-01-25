package database

import (
	"fmt"
	"log"
	"os"

	"github.com/matheuspsantos/purchase-wex/src/core/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	strCon := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	log.Println("Oppening connection with database...")
	DB, err = gorm.Open(postgres.Open(strCon))
	if err != nil {
		log.Panic("[ERROR] Database connetion fail.", err)
	}
	log.Println("Migrating entities...")
	DB.AutoMigrate(&models.Purchase{})
}

func CloseDatabase() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("Error retrieving SQL database: %v", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		log.Printf("Error closing database connection: %v", err)
	} else {
		log.Println("Database connection closed.")
	}
}
