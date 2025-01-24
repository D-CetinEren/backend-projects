package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "expense-tracker",
	Short: "A simple CLI application to manage expenses",
	Long:  "Expense Tracker is a CLI tool to add, delete, view, and summarize your expenses.",
}

// Execute initializes the root command and all subcommands.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
