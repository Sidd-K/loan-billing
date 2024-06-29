package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"loan-building/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/loans", handlers.CreateLoan).Methods("POST")
	r.HandleFunc("/loans/{id}/outstanding", handlers.GetOutstanding).Methods("GET")
	r.HandleFunc("/loans/{id}/delinquent", handlers.IsDelinquent).Methods("GET")
	r.HandleFunc("/loans/{id}/payment", handlers.MakePayment).Methods("POST")

	http.ListenAndServe(":8000", r)
}
