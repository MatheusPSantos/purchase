package repository

import (
	"log"

	"github.com/matheuspsantos/purchase-wex/src/core/models"
	database "github.com/matheuspsantos/purchase-wex/src/infra/configs"
)

func StoreNewPurchaseTransactionUseCase(p *models.Purchase) (**models.Purchase, error) {
	stored := &p
	if err := database.DB.Create(stored).Error; err != nil {
		log.Printf("Error while saving to database: %v", err)
		return nil, err
	}
	log.Printf("Stored value: %+v", stored)
	return stored, nil
}

func ListAllPurchaseTransactionUseCase(p *[]models.Purchase) *[]models.Purchase {
	database.DB.Find(p)
	return p
}
