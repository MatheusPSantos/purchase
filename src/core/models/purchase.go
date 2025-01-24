package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model
	Description string  `json:"description" validate:"max=50"`
	TransactionDate     string  `json:"transaction_date" validate:"nonzero"`
	Amount      float64 `json:"amount" validate:"min=0.01"`
}

func Validate(p *Purchase) error {
	if err := validator.Validate(p); err != nil {
		return err
	}
	return nil
}
