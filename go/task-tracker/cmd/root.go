package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task-tracker",
	Short: "A simple task tracking CLI application",
	Long:  `Manage and track your tasks through the command line interface.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Global flags and subcommands can be initialized here
}
