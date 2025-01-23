package database

import (
	"log"

	"github.com/matheuspsantos/purchase-wex/src/core/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	strCon := "host=localhost user=root password=root dbname=purchase port=5432 sslmod=disable"
	log.Println("Oppening connection with database...")
	DB, err = gorm.Open(postgres.Open(strCon))
	if err != nil {
		log.Panic("[ERROR] Database connetion failed. ", err)
	}
	log.Println("Migrating entities...")
	DB.AutoMigrate(&entities.Purchase{})
}
