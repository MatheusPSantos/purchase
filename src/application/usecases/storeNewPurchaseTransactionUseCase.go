package usecases

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/matheuspsantos/purchase-wex/src/core/dtos"
	"github.com/matheuspsantos/purchase-wex/src/core/models"
	"github.com/matheuspsantos/purchase-wex/src/infra/repository"
)

func StoreNewPurchaseTransactionUseCase(pyld io.ReadCloser) (**models.Purchase, error) {
	defer pyld.Close()
	var newPurchase dtos.PurchaseTrxRequestDTO

	if err := json.NewDecoder(pyld).Decode(&newPurchase); err != nil {
		return nil, fmt.Errorf("failed to decode payload: %w", err)
	}

	purchase, err := dtos.ValidatePurchaseInput(newPurchase)
	if err != nil {
		fmt.Errorf("%w", err)
		return nil, err
	}

	s, err := repository.StoreNewPurchaseTransactionUseCase(purchase)
	if err != nil {
		return nil, err
	}
	return s, nil
}
