package services

import (
	"github.com/google/uuid"
	"hexagonal-example/internal/domain"
	"hexagonal-example/internal/repositories"
	"log"
)

type bankService struct {
	databaseRepository repositories.DatabaseRepository
}

func New(databaseRepository repositories.DatabaseRepository) *bankService {
	return &bankService{
		databaseRepository: databaseRepository,
	}
}

func (srv bankService) WithdrawFromAccount(id uuid.UUID, amount float64) error {
	account, err := srv.databaseRepository.Get(id)
	if err != nil {
		return err
	}

	err = account.Withdraw(amount)
	if err != nil {
		return err
	}

	return nil
}

func (srv bankService) Balance(id uuid.UUID) (float64, error) {
	account, err := srv.databaseRepository.Get(id)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return account.Balance(), nil
}

func (srv bankService) Create(account domain.Account) (uuid.UUID, error) {
	account.Id = uuid.New()
	err := srv.databaseRepository.Save(&account)
	if err != nil {
		log.Println(err)
		return uuid.Nil, err
	}
	return account.Id, nil
}