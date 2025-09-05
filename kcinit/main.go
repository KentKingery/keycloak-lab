package main

import (
	"log/slog"
	"strconv"
	"time"
)

func main() {

	slog.Info("Keycloak initialization started")

	slog.Info("Acquiring token from Keycloak server")
	token := GetToken()
	expiry := time.Now().Add(time.Second * time.Duration(token.ExpiresIn))
	slog.Info("Token acquired")
	slog.Info("Token length: " + strconv.Itoa(len(token.AccessToken)))
	slog.Info("Token expiry: " + expiry.Format("2006-01-02 15:04:05"))

	slog.Info("Fetching realms from Keycloak server")

	realms := GetRealms(token.AccessToken)
	slog.Info("# of realms: " + strconv.Itoa(len(realms)))
	for _, realm := range realms {
		slog.Info("Realm ID: " + realm.ID)
		slog.Info("Realm Name: " + realm.Realm)
		slog.Info("Realm Display Name: " + realm.DisplayName)
	}

	slog.Info("Keycloak initialization ended")

}
