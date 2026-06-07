/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"
	"main/agent"
	"main/cmd"
)

func main() {
	if err := agent.InitDirectory(); err != nil {
		log.Fatalf("Error occured while trying to initalize config in user home dir : ", err)
	}

	cmd.Execute()
}
