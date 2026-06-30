package app

import (
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/greet", Greet)
	http.HandleFunc("/getCustomers", GetCustomer)
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
