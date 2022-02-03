package security

import (
	"net/http"
	"strings"
)

func BasicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//If no Authentication header is present the 'ok' return value will be false.
		username, password, ok := r.BasicAuth()
		if ok {
			expectedUsername := "test@test.com"
			expectedPassword := "12345678"

			usernameMatch := strings.Compare(username, expectedUsername) == 0
			passwordMatch := strings.Compare(password, expectedPassword) == 0

			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}

		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}
