package security

import (
	"encoding/json"
	"everywheretew.it/main/common"
	"fmt"
	"net/http"
	"strings"
)

func Login(w http.ResponseWriter, r *http.Request) {
	common.SetupCORS(&w, r)

	username := r.URL.Query()["username"][0]
	password := r.URL.Query()["password"][0]

	if strings.Compare(globalUsersDb[username].Password, password) == 0 {
		fmt.Println("@@@@ SUCCESS")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(globalUsersDb[username]); err != nil {
			panic(err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)
	}
}
