package store

import (
	"fmt"
	"os"
)

const rootDirName = ".brocode"

type directory struct {
	providers    string
	defaultModel string
	creds        string
}

// Store the final path of store files
var storeDir directory

func InitStore() error {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Accesing user home dir : %w", err)
	}

	configDir := userHomeDir + "/" + rootDirName

	err = os.Mkdir(configDir, os.ModeDir|os.ModePerm)
	if err != nil {
		// return fmt.Errorf("Creating .brocode in home dir : %w", err)
	}

	storeDir = directory{
		providers:    configDir + "/" + "providers",
		defaultModel: configDir + "/" + "default",
		creds:        configDir + "/" + "creds",
	}

	_, err = os.OpenFile(storeDir.providers, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("Creating providers file : %w", err)
	}

	_, err = os.OpenFile(storeDir.creds, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("Creating  default provider : %w", err)
	}

	_, err = os.OpenFile(storeDir.creds, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("Creating creds file : %w", err)
	}

	return nil
}
