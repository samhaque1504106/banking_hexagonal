package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Customers struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customers{
		{Name: "a", City: "aa", ZipCode: "1234"},
		{Name: "b", City: "bb", ZipCode: "4321"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["customer_id"]
	fmt.Fprint(w, customerID)
}
