package infra

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/matheuspsantos/purchase-wex/src/infra/configs"
	"github.com/matheuspsantos/purchase-wex/src/infra/routers"
)

func RunApplication() {
	configs.ConnectDatabase()
	sm := routers.NewRouter(mux.NewRouter())

	s := &http.Server{
		Addr:         ":9090",
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

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("Received terminate, graceful shutdown", sig)

	tc, err := context.WithTimeout(context.Background(), 30*time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	s.Shutdown(tc)
}
