package app

import (
	"banking/service"
	"encoding/json"
	"encoding/xml"
	fmt "fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Customer struct {
	Name    string `json:"fullname,omitempty"`
	City    string `json:"city,omitempty"`
	Zipcode string `json:"zip_code,omitempty"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	//customers := []Customer{
	//	{"Charles", "Davao City", "9000"},
	//	{"sab", "Davao City", "8000"},
	//	{"chaz", "Davao City", "7000"},
	//}

	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-type", "application/xml")
		xml.NewEncoder(w).Encode(customers)

	} else {
		w.Header().Add("Content-type", "application/json")
		json.NewEncoder(w).Encode(customers)

	}
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}
