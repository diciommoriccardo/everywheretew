package common

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetRequest(url string) []byte {
	log.Println("[INFO] - Calling api: " + url)
	var bearer = "Bearer " + TwitterBearerToken

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	return body
}
