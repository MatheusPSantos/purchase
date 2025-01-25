package entrypoints_rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matheuspsantos/purchase-wex/src/application/usecases"
)

func StorePurchaseTransaction(w http.ResponseWriter, r *http.Request) {
	log.Print("Storing purchase transaction")
	res, err := usecases.StoreNewPurchaseTransactionUseCase(r.Body)
	if err != nil {
		fmt.Errorf("%w", err)
		http.Error(w, "Error while storing purchase transaction: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		fmt.Errorf("%w", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func GetPurchaseTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idPurchase := vars["id"]
	cur := r.URL.Query().Get("currency")

	if idPurchase == "" {
		http.Error(w, "Missing or invalid purchase ID", http.StatusBadRequest)
		return
	}

	if cur == "" {
		http.Error(w, "Invalid currency format.", http.StatusBadRequest)
		return
	}

	log.Printf("Getting purchase transaction with id = %s and currency = %s", idPurchase, cur)
	res, err := usecases.GetPurchaseTransactionByIdUseCase(idPurchase, cur)
	if err != nil {
		log.Print(err)
		http.Error(w, "Error when try retrieve purchase transaction: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("JSON enconding error:  %v", err)
		http.Error(w, "Failed to encode response.", http.StatusBadRequest)
		return
	}
}

func ListAllPurchaseTransactions(w http.ResponseWriter, r *http.Request) {
	res := usecases.ListAllPurchaseTransactionsUseCase()

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("JSON encoding error: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
