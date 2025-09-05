package main

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Token struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not_before_policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

func GetToken() Token {
	var token Token

	tokenUrl := "http://localhost:8180/realms/master/protocol/openid-connect/token"
	encoding := "application/x-www-form-urlencoded"

	formData := url.Values{}
	formData.Set("client_id", "admin-cli")
	formData.Set("grant_type", "password")
	formData.Set("username", "admin")
	formData.Set("password", "p@ssw0rd")

	req, err := http.NewRequest("POST", tokenUrl, strings.NewReader(formData.Encode()))
	req.Header.Set("Content-Type", encoding)

	response, err := http.DefaultClient.Do(req)

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		slog.Error(err.Error())
	}

	err = json.Unmarshal(responseData, &token)

	return token
}
