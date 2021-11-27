package repositories

import (
	"errors"
	"github.com/google/uuid"
	"hexagonal-example/internal/domain"
)

type DatabaseRepository interface {
	Get(accountId uuid.UUID) (*domain.Account, error)
	Save(account *domain.Account) error
}

type memoryDb struct {
	instance map[uuid.UUID]*domain.Account
}

func NewMemoryDb() *memoryDb {
	memDb := make(map[uuid.UUID]*domain.Account)
	return &memoryDb{
		instance: memDb,
	}
}

func (memDb *memoryDb) Get(accountId uuid.UUID) (*domain.Account, error) {

	account := memDb.instance[accountId]

	if account == nil {
		return nil, errors.New("account doesnt exist")
	}

	return account, nil
}

func (memDb *memoryDb) Save(account *domain.Account) error {
	memDb.instance[account.Id] = account
	return nil
}