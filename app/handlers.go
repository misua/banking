package app

import (
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

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world!!")

}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Charles", "Davao City", "9000"},
		{"sab", "Davao City", "8000"},
		{"chaz", "Davao City", "7000"},
	}

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
