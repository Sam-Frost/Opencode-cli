package provider

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	providerCommand := &cobra.Command{
		Use:   "provider",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	providerCommand.AddCommand(listProvidersCommand())
	providerCommand.AddCommand(loginProviderCommand())
	return providerCommand
}

func listProvidersCommand() *cobra.Command {
	listProvidersCommand := &cobra.Command{
		Use:   "list",
		Short: "Show list of all the available providers",
		Long:  "Show list of all the available LLM providers in the world",
		Run: func(cmd *cobra.Command, args []string) {

			showAllModelFlag, _ := cmd.Flags().GetBool("all")

			if showAllModelFlag {
				fmt.Println("Showing all providers..")
			} else {
				fmt.Println("Showing only available providers...")
			}
		},
	}

	listProvidersCommand.Flags().BoolP("all", "a", false, "Shows all the providers(included unsupported ones too)")

	return listProvidersCommand
}

func loginProviderCommand() *cobra.Command {
	loginProviderCommand := &cobra.Command{
		Use:   "login",
		Short: "Show list of all the available providers",
		Long:  "Show list of all the available LLM providers in the world",
		Run: func(cmd *cobra.Command, args []string) {

			providerName, _ := cmd.Flags().GetString("provider")
			apiKey, _ := cmd.Flags().GetString("apiKey")

			if len(providerName) == 0 || len(apiKey) == 0 {
				fmt.Println("Please enter provider name and api key for the provider")
			} else {
				fmt.Printf("Provider Name : %s \n Provider Api Key : %s\n", providerName, apiKey)
			}
		},
	}

	loginProviderCommand.Flags().StringP("provider", "p", "", "Provider name for which user wants to login")
	loginProviderCommand.Flags().String("apiKey", "", "API key for the LLM provider")

	return loginProviderCommand
}
