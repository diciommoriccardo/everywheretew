package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello homepage")
}

func main() {
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":10000", router))
}
