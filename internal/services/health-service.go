package services

import (
	"hexagonal-example/internal/domain"
)

type service struct {
	Health domain.Health
}

func NewHealth(h domain.Health) *service {
	return &service{
		Health: h,
	}
}

func (healthSrv service) IsAppHealthy() bool {
	status := healthSrv.Health.CheckHealth()
	return status
}
