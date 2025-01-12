package cmd

import (
	"fmt"

	"github.com/D-CetinEren/backend-projects/go/Github-user-activity/internal"
	"github.com/spf13/cobra"
)

var activityCmd = &cobra.Command{
	Use:   "activity <username>",
	Short: "Fetch recent activity of a GitHub user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		err := internal.DisplayUserActivity(username)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(activityCmd)
}
