package dtos

type PurchaseTrxWithExchangeDTOResponse struct {
	ID                   uint64  `json:"id"`
	Description          string  `json:"description"`
	TransactionDate      string  `json:"transaction_date"`
	OriginalDollarAmount float64 `json:"amount_usd"`
	ExchangeRateUsed     float64 `json:"exchange_rate"`
	ConvertedAmount      float64 `json:"amount"`
}
