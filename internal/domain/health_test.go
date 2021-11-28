package domain

import (
	"reflect"
	"testing"
)

func TestCalculateRequestToExternalService(t *testing.T) {
	got, err := calculateRequestToExternalService()
	if err != nil {
		t.Error(err)
	}

	minWanted := 1000
	maxWanted := 10000

	if got < minWanted || got > maxWanted {
		t.Errorf("wanted external service to respond between %d and %d, got: %d", minWanted, maxWanted, got)
	}
}

func TestHealthCheckHealth(t *testing.T) {

	h := Health{}
	status := h.CheckHealth()

	wanted := "bool"
	got := reflect.TypeOf(status).String()

	if got != wanted {
		t.Errorf("wanted: %s, got: %s", wanted, got)
	}
}
