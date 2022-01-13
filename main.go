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

	//var testTweets = models.GetTweetsByUsername("google")
	//var testTwitterUser = models.GetTwitterUserByUsername("google")
	//fmt.Printf("\n######## Test Tweets: %s", testTweets)
	//fmt.Printf("\n######## Test Twitter User: %s", testTwitterUser)
	//var testTweetsByQuery = models.GetTweetsByQuerySearch("everywheretew")
	//fmt.Printf("\n######## Tweets contenenti il parametro ricercato: %d", len(testTweetsByQuery))

	log.Fatal(http.ListenAndServe(":10000", router))
}
