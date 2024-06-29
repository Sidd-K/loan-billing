package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"loan-building/models"
	"loan-building/services"
)

func CreateLoan(w http.ResponseWriter, r *http.Request) {
	var req models.CreateLoanRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	loan := services.CreateLoan(req)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(loan)
}

func GetOutstanding(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	outstanding, exists := services.GetOutstanding(id)
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]float64{"outstanding": outstanding})
}

func IsDelinquent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	delinquent, exists := services.IsDelinquent(id)
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]bool{"delinquent": delinquent})
}

func MakePayment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var payment models.Payment
	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	payment.PaidAt = time.Now()
	loan, exists := services.MakePayment(id, payment)
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(loan)
}
