package gateways

type (
	PurchaseRepository interface {
		StoreTransaction() error
	}
)
