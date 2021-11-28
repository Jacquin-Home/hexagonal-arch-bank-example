package repositories

import (
	"errors"
	"github.com/google/uuid"
	"hexagonal-example/internal/domain"
)

type memoryDb struct {
	instance map[uuid.UUID]interface{}
}

func NewMemoryDb() *memoryDb {
	memDb := make(map[uuid.UUID]interface{})
	return &memoryDb{
		instance: memDb,
	}
}

func (memDb *memoryDb) GetAccount(accountId uuid.UUID) (*domain.Account, error) {

	account := memDb.instance[accountId]
	acc := domain.Account{
		Id:    uuid.New(),
		Money: 100,
	}

	if account == nil {
		return nil, errors.New("account doesnt exist")
	}

	//return account, nil
	return &acc, nil
}

func (memDb *memoryDb) SaveAccount(account *domain.Account) error {
	memDb.instance[uuid.New()] = account
	return nil
}

func (memDb *memoryDb) SavePayment(payment domain.Payment) error {
	memDb.instance[uuid.New()] = payment

	return nil
}
