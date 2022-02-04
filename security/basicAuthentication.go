package security

import (
	"fmt"
	"net/http"
)

func BasicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//If no Authentication header is present the 'ok' return value will be false.
		username, password, ok := r.BasicAuth()
		if ok {
			//userMatch := common.GlobalUsersDb[username]
			//passwordMatch := strings.Compare(password, userMatch.Password) == 0

			fmt.Println(username)
			fmt.Println(password)
			passwordMatch := true

			if passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}

		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}
