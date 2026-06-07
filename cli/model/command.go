package model

import (
	"fmt"

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
		Run: func(cmd *cobra.Command, args []string) {
			modelName := args[0]

			if len(modelName) == 0 {
				fmt.Println("Model name is missing...")
			} else {
				fmt.Printf("%s set as default model\n", modelName)
			}
		},
	}

	return setModelCommand
}
