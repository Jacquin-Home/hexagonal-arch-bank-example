package main

import (
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"hexagonal-example/internal/domain"
	"hexagonal-example/internal/repositories"
	"hexagonal-example/internal/services"
	"log"
)

func main() {
	//dbRepository := repositories.NewMemoryDb()
	dbRepository := repositories.NewSqliteDB()

	id := uuid.New()
	acc := domain.Account{
		Id:    id,
		Money: 100,
	}

	//err := dbRepository.saveAccount(&acc)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	srv := services.NewAccountService(dbRepository)
	_, err := srv.Create(acc)
	if err != nil {
		log.Println(err)
	}
	err = srv.WithdrawFromAccount(id, 50)
	if err != nil {
		log.Println(err)
		return
	}

	//balance, err := srv.Balance(acc.Id)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//fmt.Println(balance)

}
