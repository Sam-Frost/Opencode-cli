package agent

import (
	"fmt"
	"os"
)

func InitDirectory() error {
	fmt.Println("Running init directory")

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Accesing user home dir : %w", err)
	}

	configDir := userHomeDir + "/.brocode"

	err = os.Mkdir(configDir, os.ModeDir|os.ModePerm)
	if err != nil {
		// return fmt.Errorf("Creating .brocode in home dir : %w", err)
	}

	_, err = os.OpenFile(fmt.Sprintf("%s/providers", configDir), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("Creating providers file : %w", err)
	}

	_, err = os.OpenFile(fmt.Sprintf("%s/creds", configDir), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("Creating creds file : %w", err)

	}

	_, err = os.OpenFile(fmt.Sprintf("%s/default", configDir), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("Creating creds file : %w", err)

	}

	return nil
}
