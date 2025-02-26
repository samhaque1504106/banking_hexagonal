package app

import (
	"github.com/gorilla/mux"
	"github.com/samhaque1504106/banking_hexagonal/domain"
	"github.com/samhaque1504106/banking_hexagonal/service"
	"log"
	"net/http"
)

func Start() {
	//mux := http.NewServeMux()
	router := mux.NewRouter()

	//wirings
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	//define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	//router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	//router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	//router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
