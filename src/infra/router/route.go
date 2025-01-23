package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	middleware "github.com/matheuspsantos/purchase-wex/src/infra"
	entrypoints_rest "github.com/matheuspsantos/purchase-wex/src/infra/entrypoints/rest"
)

func NewRouter(route *mux.Router) *mux.Router {
	log.Println("Registering routes...")
	sm := route

	sm.Use(middleware.SetHeaderContentTypeJson)

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/purchase", entrypoints_rest.GetPurchaseTransaction)

	postRotuer := sm.Methods(http.MethodPost).Subrouter()
	postRotuer.HandleFunc("/purchase", entrypoints_rest.StorePurchaseTransaction)

	return sm
}
