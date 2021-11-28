package domain

import (
	"log"
	"math/rand"
	"time"
)

type Health struct{}

func (h Health) CheckHealth() bool {
	// check request time
	requestTime, err := calculateRequestToExternalService()
	if err != nil {
		log.Println(err)
	}
	if requestTime >= 5000 {
		return false
	}
	return true
}

func calculateRequestToExternalService() (int, error) {
	min := 1000
	max := 10000
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(max-min) + min
	return result, nil
}
