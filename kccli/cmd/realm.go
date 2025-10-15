/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"kccli/utility"
)

type Realm struct {
	ID          string `json:"id"`
	Realm       string `json:"realm"`
	DisplayName string `json:"displayName"`
}

type NewRealm struct {
	Realm       string `json:"realm"`
	DisplayName string `json:"displayName"`
	Enabled     bool   `json:"enabled"`
}

// realmCmd represents the realm command
var realmCmd = &cobra.Command{
	Use:   "realm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		realmName := ""
		realmDisplayName := ""
		flag := ""
		if len(args) > 0 {
			flag = args[0]
		} else {
			flag = ""
		}

		switch flag {

		case "ls":
			token := utility.GetToken()
			realms := GetRealms(token.AccessToken)
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
			fmt.Fprintln(w, "ID\tNAME\tDISPLAY NAME")
			for _, realm := range realms {
				fmt.Fprintln(w, realm.ID+"\t"+realm.Realm+"\t"+realm.DisplayName)
			}
			w.Flush()

		case "rm":
			realmName = args[1]
			token := utility.GetToken()
			DeleteRealm(token.AccessToken, realmName)

		case "create":
			realmName = args[1]
			if len(args) > 2 {
				realmDisplayName = args[2]
			}
			token := utility.GetToken()
			CreateRealm(token.AccessToken, realmName, realmDisplayName)

		default:
			helpText := `
Usage:  kccli realm COMMAND

Manage realms

Commands:
  create      Create a realm
  ls          List realms
  rm          Remove realm
  `
			fmt.Println(helpText)
		}
	},
}

func init() {
	rootCmd.AddCommand(realmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// realmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// realmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func GetRealms(accessToken string) []Realm {
	var realms []Realm
	bearer := "Bearer " + accessToken

	realmURL := "http://localhost:8180/admin/realms"

	req, err := http.NewRequest("GET", realmURL, nil)
	req.Header.Add("Authorization", bearer)

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	jsonData := []byte(responseData)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(jsonData, &realms)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	return realms
}

func CreateRealm(accessToken string, realmName string, displayName string) {
	bearer := "Bearer " + accessToken
	newRealm := NewRealm{Realm: realmName, DisplayName: displayName, Enabled: true}

	jsonData, err := json.Marshal(newRealm)
	if err != nil {
		// Handle error
	}

	realmURL := "http://localhost:8180/admin/realms"

	req, err := http.NewRequest("POST", realmURL, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()

	/*
		responseData, err := io.ReadAll(response.Body)
		jsonData := []byte(responseData)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}

		err = json.Unmarshal(jsonData, &realms)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
	*/
}

func DeleteRealm(accessToken string, realmName string) {
	bearer := "Bearer " + accessToken

	realmURL := "http://localhost:8180/admin/realms/" + realmName

	req, err := http.NewRequest("DELETE", realmURL, nil)
	req.Header.Add("Authorization", bearer)

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()

	/*
		responseData, err := io.ReadAll(response.Body)
		jsonData := []byte(responseData)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}

		err = json.Unmarshal(jsonData, &realms)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
	*/
}

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
		//
		CreateUser(token.AccessToken, "globalco", "kentkingery", "Kent", "Kingery", "kent@globalco.io")
		slog.Info("Keycloak initialization ended")
*/
