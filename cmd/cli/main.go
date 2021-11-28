package main

import (
	"fmt"
	"github.com/google/uuid"
	"hexagonal-example/internal/domain"
	"hexagonal-example/internal/repositories"
	"hexagonal-example/internal/services"
	"log"
)

func main() {
	dbRepository := repositories.NewMemoryDb()

	acc := domain.Account{
		Id:    uuid.New(),
		Money: 100,
	}

	err := dbRepository.SaveAccount(&acc)
	if err != nil {
		log.Println(err)
		return
	}

	srv := services.NewAccountService(dbRepository)
	err = srv.WithdrawFromAccount(uuid.New(), 50)
	if err != nil {
		log.Println(err)
		return
	}

	balance, err := srv.Balance(acc.Id)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(balance)

}
