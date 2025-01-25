package infra

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/matheuspsantos/purchase-wex/src/infra/database"
	"github.com/matheuspsantos/purchase-wex/src/infra/routers"
)

func RunApplication() {
	database.ConnectDatabase()
	sm := routers.NewRouter(mux.NewRouter())

	s := &http.Server{
		Addr:         ":8888",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill)

	sig := <-sigChan
	log.Println("Received terminate, graceful shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.Shutdown(tc); err != nil {
		log.Printf("Error shutting down server: %v", err)
	}

	database.CloseDatabase()
	log.Println("Application shutdown complete.")
}
