package main

import (
	"log"
	"main/cmd"
	"main/store"
)

func main() {
	if err := store.InitStore(); err != nil {
		log.Fatalf("Initializing Store : %s", err)
	}

	GetProvidersListFromApi()
	cmd.Execute()
}
