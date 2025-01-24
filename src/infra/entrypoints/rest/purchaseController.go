package entrypoints_rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/matheuspsantos/purchase-wex/src/application/usecases"
)

func StorePurchaseTransaction(w http.ResponseWriter, r *http.Request) {
	log.Print("Storing purchase transaction")
	res, err := usecases.StoreNewPurchaseTransactionUseCase(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error while storing purchase transaction", http.StatusBadRequest)
		return
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("JSON encoding error: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func GetPurchaseTransaction(w http.ResponseWriter, r *http.Request) {
	log.Print("Getting purchase transaction")
}

func ListAllPurchaseTransactions(w http.ResponseWriter, r *http.Request) {
	res := usecases.ListAllPurchaseTransactionsUseCase()

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("JSON encoding error: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
