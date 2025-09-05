package main

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
)

type Realm struct {
	ID          string `json:"id"`
	Realm       string `json:"realm"`
	DisplayName string `json:"displayName"`
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
