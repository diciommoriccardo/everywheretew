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

func getOrderData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello webHook")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/order", getOrderData)
	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	ConnectToCluster()
	fmt.Println("Rest API v2.0 - Automated Decision")
	handleRequests()
}
