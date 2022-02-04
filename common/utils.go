package common

import (
	"net/http"
)

func StringContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func HandleOptions(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
