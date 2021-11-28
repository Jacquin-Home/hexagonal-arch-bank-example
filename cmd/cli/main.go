package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"hexagonal-example/internal/domain"
	"hexagonal-example/internal/repositories"
	"hexagonal-example/internal/services"
	"log"
)

func main() {
	//dbRepository := repositories.NewMemoryDb()
	dbRepository := repositories.NewSqliteDB()

	acc := domain.Account{
		Money: 100,
	}

	srv := services.NewAccountService(dbRepository)
	id, err := srv.Create(acc)
	if err != nil {
		log.Println(err)
	}
	err = srv.WithdrawFromAccount(id, 50)
	if err != nil {
		log.Println(err)
		return
	}

	balance, err := srv.Balance(id)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(balance)

}
