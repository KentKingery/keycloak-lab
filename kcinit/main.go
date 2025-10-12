package main

import (
	"kcinit/cmd"
)

//"fmt"
//"log/slog"
//"strconv"
//"time"

//"github.com/spf13/cobra"

func main() {

	cmd.Execute()

	/*
		slog.Info("Keycloak initialization started")

		slog.Info("Acquiring token from Keycloak server")
		token := GetToken()
		expiry := time.Now().Add(time.Second * time.Duration(token.ExpiresIn))
		slog.Info("Token acquired")
		slog.Info("Token length: " + strconv.Itoa(len(token.AccessToken)))
		slog.Info("Token expiry: " + expiry.Format("2006-01-02 15:04:05"))

		slog.Info("Fetching realms from Keycloak server")

		realms := GetRealms(token.AccessToken)
		fmt.Println("# of realms: " + strconv.Itoa(len(realms)))
		fmt.Println("ID\t\t\t\t\tName\t\tDisplay Name")
		for _, realm := range realms {
			fmt.Println(realm.ID + "\t" + realm.Realm + "\t\t" + realm.DisplayName)
		}

		CreateRealm(token.AccessToken, "globalco", "GlobalCo")
		// DeleteRealm(token.AccessToken, "globalco")
		CreateUser(token.AccessToken, "globalco", "kentkingery", "Kent", "Kingery", "kent@globalco.io")
		slog.Info("Keycloak initialization ended")
	*/
}
