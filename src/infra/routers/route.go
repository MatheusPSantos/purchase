package routers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	entrypoints_rest "github.com/matheuspsantos/purchase-wex/src/infra/entrypoints/rest"
	"github.com/matheuspsantos/purchase-wex/src/infra/middlewares"
)

func NewRouter(route *mux.Router) *mux.Router {
	log.Println("Registering routes...")
	sm := route

	sm.Use(middlewares.SetHeaderContentTypeJson)

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/purchase/{id}", entrypoints_rest.GetPurchaseTransaction)
	getRouter.HandleFunc("/purchase", entrypoints_rest.ListAllPurchaseTransactions)

	postRotuer := sm.Methods(http.MethodPost).Subrouter()
	postRotuer.HandleFunc("/purchase", entrypoints_rest.StorePurchaseTransaction)

	return sm
}
