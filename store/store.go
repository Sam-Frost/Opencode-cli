package store

import (
	"fmt"
	"os"
)

func SetDefaultModel(modelName string) error {
	file, err := os.Create(storeDir.defaultModel)
	if err != nil {
		return fmt.Errorf("Creating default model file : %w", err)
	}

	_, err = file.Write([]byte(modelName))

	if err != nil {
		return fmt.Errorf("Writing to default model file : %w", err)
	}
	return nil
}

func GetDefaultModel() (string, error) {

	file, err := os.Open(storeDir.defaultModel)

	if err != nil {
		return "", fmt.Errorf("Opening default model file : %w", err)
	}

	fileBytes := make([]byte, 30)

	_, err = file.Read(fileBytes)

	if err != nil {
		return "", fmt.Errorf("Reading default model from file : %w ", err)
	}

	data := string(fileBytes)

	return data, nil
}
