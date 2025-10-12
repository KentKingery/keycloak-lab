package main

import (
	"kccli/cmd"
	//"log/slog"
	//"os"
	//"strconv"
)

func main() {

	// slog.Info("kccli - Keycloak Command Line Interface")

	// slog.Info(strconv.Itoa(len(os.Args[1:])))

	cmd.Execute()

	/*
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
	*/
}
