package services

import (
	"github.com/google/uuid"
	"hexagonal-example/internal/domain"
	"hexagonal-example/internal/repositories"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	memDb := repositories.NewMemoryDb()
	srv := New(memDb)

	srvInstance := &bankService{}

	if reflect.TypeOf(srv) != reflect.TypeOf(srvInstance) {
		t.Errorf("wanted: instance of *services, got: instance of %v", reflect.TypeOf(srvInstance))
	}
}

func TestServiceWithdrawFromAccount(t *testing.T) {

	memDb := repositories.NewMemoryDb()
	srv := New(memDb)

	id := uuid.New()
	acc := domain.Account{
		Id: id,
		Money: 100,
	}

	err := srv.databaseRepository.Save(&acc)
	if err != nil {
		t.Error(err)
	}

	err = srv.WithdrawFromAccount(id, 100)
	if err != nil {
		t.Error(err)
	}

	if acc.Balance() != 0 {
		t.Error("amount should be zero")
	}
}

func TestServiceBalance(t *testing.T) {

	memDb := repositories.NewMemoryDb()
	srv := New(memDb)

	wanted := 100.1

	id := uuid.New()
	acc := domain.Account{
		Id: id,
		Money: wanted,
	}

	err := srv.databaseRepository.Save(&acc)
	if err != nil {
		t.Error(err)
	}

	balance, err := srv.Balance(id)
	if err != nil {
		t.Error(err)
	}

	if wanted != balance {
		t.Errorf("wanted: %T, got: %T", wanted, balance)
	}
}