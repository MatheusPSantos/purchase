package usecases

import (
	"github.com/matheuspsantos/purchase-wex/src/core/models"
	"github.com/matheuspsantos/purchase-wex/src/infra/repository"
	"gorm.io/gorm"
)

func ListAllPurchaseTransactionsUseCase() *gorm.DB {
	var p []models.Purchase
	return repository.ListAllPurchaseTransactionUseCase(&p)
}
