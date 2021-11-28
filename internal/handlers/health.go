package handlers

import (
	"encoding/json"
	"hexagonal-example/internal/services"
	"net/http"
)

type HTTPHealthHandler struct {
	HealthService services.InterfaceHealth
}

func (h HTTPHealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	status := h.HealthService.IsAppHealthy()

	json.NewEncoder(w).Encode(status)
}
