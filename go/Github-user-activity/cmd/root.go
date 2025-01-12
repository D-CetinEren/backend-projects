package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Github-user-activity",
	Short: "Fetch and display GitHub user activity",
	Long:  "A CLI tool to fetch and display recent activity of a GitHub user using the GitHub API.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use the 'activity' command to fetch GitHub user activity.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
