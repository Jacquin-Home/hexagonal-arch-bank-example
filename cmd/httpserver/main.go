package main

import (
	"github.com/gorilla/mux"
	"hexagonal-example/internal/domain"
	"hexagonal-example/internal/handlers"
	"hexagonal-example/internal/repositories"
	"hexagonal-example/internal/services"
	"log"
	"net/http"
)

func main() {

	// database instance
	dbRepository := repositories.NewMemoryDb()
	// instantiate database service
	srv := services.New(dbRepository)

	// bank http handler
	bankHandler := handlers.NewHTTPHandler(srv)

	// health service
	healthService := services.NewHealth(domain.Health{})
	// health http handler
	healthHandler := handlers.HTTPHealthHandler{
		HealthService: healthService,
	}

	r := mux.NewRouter()
	r.HandleFunc("/account/create", bankHandler.Create).Methods(http.MethodPost)
	r.HandleFunc("/account/balance/{id}", bankHandler.Balance).Methods(http.MethodGet)
	r.HandleFunc("/health", healthHandler.HealthCheck).Methods(http.MethodGet)
	log.Println("listening at port 7000...")
	panic(http.ListenAndServe(":7000", r))

}
