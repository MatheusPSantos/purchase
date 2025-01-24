package usecases

import (
	"encoding/json"
	"io"
	"log"

	"github.com/matheuspsantos/purchase-wex/src/core/models"
	"github.com/matheuspsantos/purchase-wex/src/infra/repository"
)

func StoreNewPurchaseTransactionUseCase(pyld io.ReadCloser) (**models.Purchase, error) {
	var newPurchase models.Purchase
	json.NewDecoder(pyld).Decode(&newPurchase)
	log.Printf("%#v", newPurchase)
	if err := models.Validate(&newPurchase); err != nil {
		return nil, err
	}
	s, err := repository.StoreNewPurchaseTransactionUseCase(&newPurchase)
	if err != nil {
		return nil, err
	}
	return s, nil
}
