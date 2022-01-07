package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello homepage")
}

func Router() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/api/order", GetOrderData).Methods("POST")
	router.HandleFunc("/api/products", GetProductData).Methods("POST")
	router.HandleFunc("/api/cluster/{id}", getCluster).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	fmt.Println("EverywhereTew - Automated Decision")
	Router()
}
