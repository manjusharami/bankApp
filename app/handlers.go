package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	CustomerId            int     `json:"customerID" xml:"cutomerId"`
	CustomerName          string  `json:"firstName" xml:"name"`
	CustomerAccountNumber int     `json:"accountNumber" xml:"accNumber"`
	Balance               float64 `json:"balance" xml:"bal"`
}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to chinnu world")
}
func GetCustomer(w http.ResponseWriter, r *http.Request) {
	getCustomerData := []Customer{
		{
			CustomerId:            10,
			CustomerName:          "manjusha",
			CustomerAccountNumber: 33999,
			Balance:               10000.99,
		},
		{
			CustomerId:            20,
			CustomerName:          "chinnu",
			CustomerAccountNumber: 33999,
			Balance:               10000.00,
		},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(getCustomerData)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(getCustomerData)
	}

}
