package main

import (
	"fmt"
	"io"
	"net/http"
)

func GetProvidersListFromApi() (string, error) {
	fmt.Println("Function got called?")
	resp, err := http.Get("https://models.dev/api.json")
	if err != nil {
		return "", fmt.Errorf("Hitting GET API : %w", err)
	}

	body, err := io.ReadAll(resp.Body)

	fmt.Println(string(body))

	return "", nil
}
