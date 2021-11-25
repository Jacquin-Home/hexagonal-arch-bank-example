package bank_handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"hexagonal-example/internal/core/domain/bank"
	"hexagonal-example/internal/core/ports"
	"io/ioutil"
	"log"
	"net/http"
)

type HTTPHandler struct {
	bankService ports.BankServiceInterface
}

func NewHTTPHandler (bankService ports.BankServiceInterface) *HTTPHandler {
	return &HTTPHandler{
		bankService: bankService,
	}
}

func (h *HTTPHandler) SendSMS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("smsSent")
	w.Write([]byte("ok"))
}

func (h *HTTPHandler) Balance(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	pathVars := mux.Vars(r)
	id := pathVars["id"]
	accUuid := uuid.MustParse(id)
	balance, err := h.bankService.Balance(accUuid)
	if err != nil {
		log.Println(err)
	}

	jsonData := map[string]interface{}{
		"id": accUuid,
		"balance": balance,
	}

	err = json.NewEncoder(w).Encode(jsonData)
	if err != nil {
		log.Println(err)
	}
}

func (h *HTTPHandler) Create(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	var acc bank.Account
	err = json.Unmarshal(body, &acc)
	if err != nil {
		log.Println(err)
	}

	newId, err := h.bankService.Create(acc)
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
