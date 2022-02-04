package security

import (
	"encoding/json"
	"everywheretew.it/main/common"
	"everywheretew.it/main/models"
	"fmt"
	"net/http"
	"strings"
)

func Login(w http.ResponseWriter, r *http.Request) {
	common.SetupCORS(&w, r)

	// in memory user db
	usersDb := make(map[string]models.SystemUser)
	usersDb["test@utente.com"] = models.SystemUser{
		Id:       1,
		Email:    "test@utente.com",
		Name:     "Test",
		Role:     "Utente",
		Password: "Test_001",
	}

	usersDb["test@admin.com"] = models.SystemUser{
		Id:       2,
		Email:    "test@admin.com",
		Name:     "Test",
		Role:     "Admin",
		Password: "Test_001",
	}

	username := r.URL.Query()["username"][0]
	password := r.URL.Query()["password"][0]

	if strings.Compare(usersDb[username].Password, password) == 0 {
		fmt.Println("@@@@ SUCCESS")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(usersDb[username]); err != nil {
			panic(err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)
	}
}
