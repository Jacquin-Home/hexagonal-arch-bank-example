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
	accService := services.NewAccountService(dbRepository)

	// bank http handler
	accountHandler := handlers.NewHTTPHandler(accService)

	// health service
	healthService := services.NewHealth(domain.Health{})
	// health http handler
	healthHandler := handlers.HTTPHealthHandler{
		HealthService: healthService,
	}

	payService := services.NewPayment(dbRepository)
	paymentHandler := handlers.NewHTTPPaymentHandler(payService)

	r := mux.NewRouter()
	r.HandleFunc("/account/create", accountHandler.Create).Methods(http.MethodPost)
	r.HandleFunc("/account/balance/{id}", accountHandler.Balance).Methods(http.MethodGet)
	r.HandleFunc("/payment", paymentHandler.RegisterPayment).Methods(http.MethodPost)
	r.HandleFunc("/health", healthHandler.HealthCheck).Methods(http.MethodGet)
	log.Println("listening at port 7000...")
	panic(http.ListenAndServe(":7000", r))

}
