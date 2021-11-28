package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"hexagonal-example/internal/domain"
	"io/ioutil"
	"log"
	"net/http"
)

type AccountServiceInterface interface {
	Balance(id uuid.UUID) (float64, error)
	Create(domain.Account) (uuid.UUID, error)
}

type HTTPAccountHandler struct {
	accountService AccountServiceInterface
}

func NewHTTPHandler(accountService AccountServiceInterface) *HTTPAccountHandler {
	return &HTTPAccountHandler{
		accountService: accountService,
	}
}

func (h *HTTPAccountHandler) SendSMS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("smsSent")
	w.Write([]byte("ok"))
}

func (h *HTTPAccountHandler) Balance(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	pathVars := mux.Vars(r)
	id := pathVars["id"]
	accUuid := uuid.MustParse(id)
	balance, err := h.accountService.Balance(accUuid)
	if err != nil {
		log.Println(err)
	}

	jsonData := map[string]interface{}{
		"id":      accUuid,
		"balance": balance,
	}

	err = json.NewEncoder(w).Encode(jsonData)
	if err != nil {
		log.Println(err)
	}
}

func (h *HTTPAccountHandler) Create(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	var acc domain.Account
	err = json.Unmarshal(body, &acc)
	if err != nil {
		log.Println(err)
	}

	newId, err := h.accountService.Create(acc)
	if err != nil {
		log.Println(err)
	}

	ret := map[string]string{
		"id": newId.String(),
	}

	err = json.NewEncoder(w).Encode(ret)
	if err != nil {
		log.Println(err)
	}
}
