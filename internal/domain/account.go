package domain

import (
	"errors"
	"github.com/google/uuid"
)

type Account struct {
	Id    uuid.UUID
	Money float64
}

func (acc Account) Balance() (balance float64) {
	return acc.Money
}

func (acc *Account) Withdraw(amount float64) error {
	if (acc.Money - amount) < 0 {
		return errors.New("no enough funds")
	}

	acc.Money = acc.Balance() - amount

	return nil
}
