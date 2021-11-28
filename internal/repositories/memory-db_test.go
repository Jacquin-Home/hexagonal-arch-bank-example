package repositories

import (
	"hexagonal-example/internal/domain"
	"reflect"
	"testing"
)

func TestNewMemoryDb(t *testing.T) {

	memDb := NewMemoryDb()

	mockedMemDb := make(map[int]*domain.Account)

	if !reflect.DeepEqual(memDb.account, mockedMemDb) {
		t.Errorf("wanted: %v, got: %v", memDb.account, mockedMemDb)
	}
}

func TestMemoryDbGetAccountDoesntExist(t *testing.T) {
	memDb := NewMemoryDb()
	_, err := memDb.Get(0)
	if err.Error() != "account doesnt exist" {
		t.Error(err)
	}
}

func TestMemoryDbGetAccountExists(t *testing.T) {
	memDb := NewMemoryDb()

	wanted := domain.Account{
		Id:    1,
		Money: 100,
	}

	err := memDb.Save(&wanted)
	if err != nil {
		t.Error(err)
	}

	got, err := memDb.Get(1)
	if err != nil {
		t.Error(err)
	}

	if got.Id != wanted.Id {
		t.Errorf("wanted: %d, got: %d", wanted.Id, got.Id)
	}
}
