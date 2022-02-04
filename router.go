package main

import (
	"everywheretew.it/main/security"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler

		if strings.Compare(route.Name, "Login") == 0 || strings.Compare(route.Method, "OPTIONS") == 0 {
			handler = route.HandlerFunc
		} else {
			handler = security.BasicAuth(route.HandlerFunc)
		}
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
