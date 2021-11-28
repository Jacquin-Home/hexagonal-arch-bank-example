package services

import (
	"hexagonal-example/internal/domain"
	"log"
)

type PaymentInterface interface {
	SavePayment(payment domain.Payment) error
}

type PaymentService struct {
	PaymentRepository PaymentInterface
}

func NewPayment(paymentRepository PaymentInterface) *PaymentService {
	return &PaymentService{
		PaymentRepository: paymentRepository,
	}
}

func (srv PaymentService) DoPayment(payment domain.Payment) (bool, error) {
	err := payment.Pay()
	if err != nil {
		log.Println(err)
		return false, err
	}

	err = validatePayment(payment)
	if err != nil {
		log.Println(err)
		return false, err
	}

	err = srv.PaymentRepository.SavePayment(payment)
	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, nil
}

func validatePayment(payment domain.Payment) error {
	// check if payment is legit
	log.Printf("validating payment id: %s\n", payment.Id)
	log.Println("valid!")
	return nil
}
