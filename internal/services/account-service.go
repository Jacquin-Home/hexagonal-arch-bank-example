package services

import (
	"github.com/google/uuid"
	"hexagonal-example/internal/domain"
	"log"
)

type AccountRepositoryInterface interface {
	GetAccount(accountId uuid.UUID) (*domain.Account, error)
	SaveAccount(account *domain.Account) error
}

type accountService struct {
	accountRepository AccountRepositoryInterface
}

func NewAccountService(databaseRepository AccountRepositoryInterface) *accountService {
	return &accountService{
		accountRepository: databaseRepository,
	}
}

func (srv accountService) WithdrawFromAccount(id uuid.UUID, amount float64) error {
	account, err := srv.accountRepository.GetAccount(id)
	if err != nil {
		return err
	}

	err = account.Withdraw(amount)
	if err != nil {
		return err
	}

	return nil
}

func (srv accountService) Balance(id uuid.UUID) (float64, error) {
	account, err := srv.accountRepository.GetAccount(id)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return account.Balance(), nil
}

func (srv accountService) Create(account domain.Account) (uuid.UUID, error) {
	account.Id = uuid.New()
	err := srv.accountRepository.SaveAccount(&account)
	if err != nil {
		log.Println(err)
		return uuid.Nil, err
	}
	return account.Id, nil
}
