package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/samhaque1504106/banking_hexagonal/service"
	"net/http"
)

type Customers struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomer(status)
	if err != nil {
		writeResponse(w, err.Code, err.Message)
	} else {
		writeResponse(w, http.StatusOK, customers)
	}

	//w.Header().Add("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(customers)
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	fmt.Println("Requested Customer ID:", id)
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		//w.Header().Add("Content-Type", "application/json")
		//w.WriteHeader(err.Code)
		//json.NewEncoder(w).Encode(err.AsMessage())

	} else {
		writeResponse(w, http.StatusOK, customer)
		//w.Header().Add("Content-Type", "application/json")
		//w.WriteHeader(http.StatusOK)
		//json.NewEncoder(w).Encode(customer)
	}
}
func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
