package agent

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	agentCommand := &cobra.Command{
		Use:   "agent",
		Short: "This is short description",
		Long:  "This is the longer one",
		Run: func(cmd *cobra.Command, args []string) {
			userPrompt, _ := cmd.Flags().GetString("prompt")

			fmt.Println("User Prompt : ", userPrompt)
		},
	}

	agentCommand.Flags().StringP("prompt", "p", "", "User prompt for the LLM")

	return agentCommand
}
