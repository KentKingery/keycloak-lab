package main

import (
	"log/slog"
	"net/http"
)

func main() {
	slog.Info("Keycloak Client")
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://api.example.com/data", nil) // For GET, body is nil
	if err != nil {
		// Handle error
	}
	resp, err := client.Do(req)
	if err != nil {

	}
	defer resp.Body.Close()
}
