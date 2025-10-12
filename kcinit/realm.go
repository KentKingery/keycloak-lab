package main

import (
	"bytes"
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

type NewRealm struct {
	Realm       string `json:"realm"`
	DisplayName string `json:"displayName"`
	Enabled     bool   `json:"enabled"`
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
