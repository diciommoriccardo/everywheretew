package main

import (
	"fmt"
	"net/http"
)

func GetOrderData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello webHook")
}

func GetProductData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello webHook")
}
