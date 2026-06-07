package agent

import (
	"fmt"
	"os"
)

func InitDirectory() error {
	fmt.Println("Running init directory")

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error accessing user home directory")
	}

	configDir := userHomeDir + "/.brocode"

	err = os.Mkdir(configDir, os.ModeDir|os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	_, err = os.OpenFile(fmt.Sprintf("%s/providers", configDir), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Error occured while creating providers file : %w", err)
	}

	// _, err = os.OpenFile(fmt.Sprintf("%s/creds", configDir), os.O_RDWR|os.O_CREATE, 0666)
	// if err != nil {
	// 	fmt.Println("Eror")
	// }

	// _, err = os.OpenFile(fmt.Sprintf("%s/default", configDir), os.O_RDWR|os.O_CREATE, 0666)
	// if err != nil {
	// 	fmt.Println("Eror")
	// }

	return nil
}
