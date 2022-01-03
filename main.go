package main

import (
	"fmt"
	"log"
	"net/http"

	//client "github.com/diciommoriccardo/everywheretew/mongo"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello homepage")
}

func getOrderData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello webHook")
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/order", getOrderData)
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Automated Decision")
	handleRequests()
}
