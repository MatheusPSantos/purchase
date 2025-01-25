package usecases

import (
	"github.com/matheuspsantos/purchase-wex/src/core/models"
	"github.com/matheuspsantos/purchase-wex/src/infra/repository"
)

func ListAllPurchaseTransactionsUseCase() *[]models.Purchase {
	var p []models.Purchase
	return repository.ListAllPurchaseTransactionUseCase(&p)
}
