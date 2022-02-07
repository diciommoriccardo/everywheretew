package security

import "everywheretew.it/main/models"

var globalUsersDb = map[string]models.SystemUser{
	"test@utente.com": {
		Id:       1,
		Email:    "test@utente.com",
		Name:     "Test",
		Role:     "Utente",
		Password: "Test_001",
	},
	"test@admin.com": {
		Id:       2,
		Email:    "test@admin.com",
		Name:     "Test",
		Role:     "Admin",
		Password: "Test_001",
	},
}
