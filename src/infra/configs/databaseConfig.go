package configs

import (
	"log"

	"github.com/matheuspsantos/purchase-wex/src/core/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	strCon := "host=172.23.0.2 user=wex password=wex dbname=wex port=5432 sslmode=disable"

	log.Println("Oppening connection with database...")
	DB, err = gorm.Open(postgres.Open(strCon))
	if err != nil {
		log.Panic("[ERROR] Database connetion fail.", err)
	}
	log.Println("Migrating entities...")
	DB.AutoMigrate(&models.Purchase{})
}
