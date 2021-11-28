package handlers

import (
	"github.com/google/uuid"
	"hexagonal-example/internal/domain"
	"log"
	"net/http"
)

type PaymentInterface interface {
	DoPayment(payment domain.Payment) (bool, error)
}

type HTTPPaymentHandler struct {
	paymentService PaymentInterface
}

func NewHTTPPaymentHandler(paymentService PaymentInterface) *HTTPPaymentHandler {
	return &HTTPPaymentHandler{
		paymentService: paymentService,
	}
}

func (srv *HTTPPaymentHandler) RegisterPayment(w http.ResponseWriter, r *http.Request) {
	payment := domain.Payment{
		Id:    uuid.New(),
		Money: 100,
	}
	status, err := srv.paymentService.DoPayment(payment)
	if err != nil {
		log.Println(err)
		w.Write([]byte("payment error"))
		return
	}
	log.Printf("payment succeed, status: %T", status)
	w.Write([]byte("payment"))
}
