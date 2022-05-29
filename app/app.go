package app

import (
	"banking/domain"
	"banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	//mux := http.NewServeMux()
	router := mux.NewRouter()

	//wiring
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	//define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	//router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	//router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
