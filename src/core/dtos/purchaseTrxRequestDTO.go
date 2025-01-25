package dtos

import (
	"errors"
	"fmt"
	"time"

	"github.com/matheuspsantos/purchase-wex/src/core/models"
)

const (
	layout = "2006-01-02"
)

type PurchaseTrxRequestDTO struct {
	Description     string  `json:"description"`
	TransactionDate string  `json:"transaction_date"`
	Amount          float64 `json:"amount"`
}

func ValidatePurchaseInput(input PurchaseTrxRequestDTO) (*models.Purchase, error) {
	if len(input.Description) > 50 {
		return nil, errors.New("description must not exceed 50 characters")
	}

	parsedDate, err := time.Parse(layout, input.TransactionDate)
	if err != nil {
		return nil, fmt.Errorf("transaction_date must be in the format YYYY-MM-DD")
	}

	if input.Amount <= 0 {
		return nil, errors.New("amount must be a positive number")
	}

	return &models.Purchase{
		Description:     input.Description,
		TransactionDate: parsedDate,
		Amount:          input.Amount,
	}, nil
}
