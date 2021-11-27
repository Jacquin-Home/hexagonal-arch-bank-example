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

	dbRepository := repositories.NewMemoryDb()

	srv := services.New(dbRepository)

	apiHandler := handlers.NewHTTPHandler(srv)

	x := domain.Health{}

	healthService := services.NewHealth(x)
	healthHandler := handlers.HTTPHealthHandler{
		HealthService: healthService,
	}

	r := mux.NewRouter()
	r.HandleFunc("/create", apiHandler.Create).Methods(http.MethodPost)
	r.HandleFunc("/balance/{id}", apiHandler.Balance).Methods(http.MethodGet)
	r.HandleFunc("/health", healthHandler.HealthCheck).Methods(http.MethodGet)
	log.Println("listening at port 7000...")
	panic(http.ListenAndServe(":7000", r))

}
