package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "expense-tracker",
		Short: "A simple CLI expense tracker application",
		Long:  "A simple CLI application to add, view, update, delete, and summarize your expenses.",
	}

	// Add subcommands
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(summaryCmd)
	rootCmd.AddCommand(deleteCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
