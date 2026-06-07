package model

import (
	"errors"
	"fmt"
	"main/store"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	modelsCommand := &cobra.Command{
		Use:   "model",
		Short: "Show the list of models available",
		Long:  "Show the list of all the models available in the world",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	modelsCommand.AddCommand(getModelsCommand())
	modelsCommand.AddCommand(setModelCommand())

	return modelsCommand
}

func getModelsCommand() *cobra.Command {

	getModelsCommand := &cobra.Command{
		Use:   "list",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

			getAllModelsFlag, _ := cmd.Flags().GetBool("all")

			if getAllModelsFlag {
				fmt.Println("get all the models")
			} else {
				fmt.Println("get available models only")
			}

			getProviderName, _ := cmd.Flags().GetString("provider")

			if len(getProviderName) > 0 {
				fmt.Printf("Showing all the models of %s :\n", getProviderName)

			}

		},
	}
	getModelsCommand.Flags().BoolP("all", "a", false, "Get list of all the models in the world")
	getModelsCommand.Flags().StringP("provider", "p", "", "Get list of all the models of the given provider")

	return getModelsCommand
}

func setModelCommand() *cobra.Command {
	setModelCommand := &cobra.Command{
		Use:   "set",
		Short: "Set the default model for brocode",
		Long:  "Set the default LLM model to be used for each user prompt by brocode",
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) == 0 || len(args[0]) == 0 {
				return errors.New("Model name is missing...")
			}

			// TODO :  Verify if the model name acutally exist

			err := store.SetDefaultModel(args[0])

			if err != nil {
				return fmt.Errorf("Saving default model : %w", err)
			}

			return nil
		},
	}

	return setModelCommand
}
