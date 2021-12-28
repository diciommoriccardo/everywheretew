package everywheretew

import (
	webHookData "everywheretew/src/routes/webHook"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", func(http.ResponseWriter, *http.Request) {})
	router.HandleFunc("/orders", webHookData.HandleOrderWebHook).Methods("POST")
	router.HandleFunc("/products", webHookData.HandleProductsWebHook).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	handleRequests()
}
