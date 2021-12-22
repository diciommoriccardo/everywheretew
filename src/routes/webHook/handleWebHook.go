package webHook

import (
	"fmt"
	"log"
	"net/http"
)

func HandleOrderWebHook() {
	handleWebHookPost()
}

func handleWebHookPost() {
	http.HandleFunc("/order", handleOrderData)
	http.HandleFunc("/product", handleProductData)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func handleOrderData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the route /order!")
	//do something with db
}

func handleProductData(w http.ResponseWriter, r *http.Request) {
	//do smt with db
}
