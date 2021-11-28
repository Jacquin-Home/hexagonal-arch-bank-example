package domain

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type Payment struct {
	Id          uuid.UUID
	FromAccount uuid.UUID
	ToAccount   uuid.UUID
	Money       float64
}

func (p *Payment) Pay() error {
	if p.Money <= 10 {
		return errors.New(
			"poor clients are not accepted, find something more expensive to buy",
		)
	}
	fmt.Println(p.Id)
	fmt.Println(p.FromAccount)
	fmt.Println(p.ToAccount)
	fmt.Println(p.Money)

	return nil
}
