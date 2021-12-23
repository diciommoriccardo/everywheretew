package everywheretew

import (
	routes "everywheretew/src/helpers"
)

/*func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}*/

//test

func main() {
	routes.SetAllRoutes()
}
