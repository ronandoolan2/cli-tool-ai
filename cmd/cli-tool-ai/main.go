// main.go
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/ronandoolan2/cli-tool-ai/pkg/actions"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "your-app",
		Short: "A CLI tool for generating code with AI",
	}

	genCmd := &cobra.Command{
		Use:   "gen <prompt>",
		Short: "Generate code with AI",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			prompt := args[0]
			err := actions.GenerateCode(prompt)
			if err != nil {
				fmt.Println("Error generating code:", err)
				os.Exit(1)
			}
		},
	}

	rootCmd.AddCommand(genCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
}
