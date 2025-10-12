package main

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
)

type User struct {
	ID        string `json:"id"`
	UserName  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type NewUser struct {
	ID            string `json:"id"`
	UserName      string `json:"username"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"emailVerified"`
	Enabled       bool   `json:"enabled"`
}

func CreateUser(accessToken string, realmName string, userName string, firstName string, lastName string, email string) {
	bearer := "Bearer " + accessToken
	newUser := NewUser{UserName: userName, FirstName: firstName, LastName: lastName, Email: email, EmailVerified: true, Enabled: true}

	jsonData, err := json.Marshal(newUser)
	if err != nil {
		slog.Error(err.Error())
	}

	userURL := "http://localhost:8180/admin/realms/" + realmName + "/users"

	req, err := http.NewRequest("POST", userURL, bytes.NewBuffer(jsonData))
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
