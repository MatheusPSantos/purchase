package entrypoints_rest

import (
	"log"
	"net/http"
)

type PurchaseController struct{}

func StorePurchaseTransaction(w http.ResponseWriter, r *http.Request) {
	log.Print("Storing purchase transaction")

}

func GetPurchaseTransaction(w http.ResponseWriter, r *http.Request) {
	log.Print("Getting purchase transaction")
}
