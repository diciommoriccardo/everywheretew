package main

import (
	"everywheretew.it/main/common"
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	common.SetupCORS(&w, r)
	fmt.Println("Hello homepage")
}

func main() {
	router := NewRouter()
	//var testTweets = models.GetTweetsByUsername("AndreaDraghetti")
	//var testTwitterUser = models.GetTwitterUserByUsername("AndreaDraghetti")
	//fmt.Printf("\n######## Test Tweets: %s", testTweets)
	//fmt.Printf("\n######## Test Twitter User: %s", testTwitterUser)

	log.Fatal(http.ListenAndServe(":10000", router))
}
