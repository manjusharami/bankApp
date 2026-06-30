package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manjusharami/bankApp/domain"
	"github.com/manjusharami/bankApp/service"
)

func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars((r))
	fmt.Fprintf(w, vars["customer_id"])
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Post the request received")

}

func Start() {
	//In Go, when we say mux, it usually means an HTTP request multiplexer. Its main job is to match incoming request URLs against a pre-defined set of routes and do something when there is a match. Simply put, a mux acts as a gateway into your application.
	router := mux.NewRouter() //multiplexer
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRespository())}
	router.HandleFunc("/getALlCustomers", ch.GetAllCustomers).Methods(http.MethodGet)

	router.HandleFunc("/createCustomer", CreateCustomer).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe("localhost:3000", router))
}
