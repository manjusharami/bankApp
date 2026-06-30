package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/manjusharami/bankApp/service"
)

type Customer struct {
	CustomerId            int     `json:"customerID" xml:"cutomerId"`
	CustomerName          string  `json:"firstName" xml:"name"`
	CustomerAccountNumber int     `json:"accountNumber" xml:"accNumber"`
	Balance               float64 `json:"balance" xml:"bal"`
}
type CustomerHandlers struct {
	service service.CustomerService
}

  func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
    customers, err := ch.service.GetAllCustomers()
    if err != nil {
        http.Error(w, "Unable to fetch customers", http.StatusInternalServerError)
        return
    }

    // Decide response format
    if r.Header.Get("Accept") == "application/xml" {
        w.Header().Set("Content-Type", "application/xml")
        xml.NewEncoder(w).Encode(customers)
        return
    }

    // Default JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(customers)
}



