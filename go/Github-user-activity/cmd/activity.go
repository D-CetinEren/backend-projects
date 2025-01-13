package cmd

import (
	"fmt"
	"os"

	"github.com/D-CetinEren/backend-projects/go/Github-user-activity/internal/api"
	"github.com/D-CetinEren/backend-projects/go/Github-user-activity/internal/filters"
	"github.com/D-CetinEren/backend-projects/go/Github-user-activity/internal/formatter"

	"github.com/spf13/cobra"
)

var eventType string

var activityCmd = &cobra.Command{
	Use:   "activity [username]",
	Short: "Fetch recent GitHub activity for a user",
	Long: `Fetch recent GitHub activity for the specified user and display it in the terminal.
You can filter activities by type (e.g., push, issue, star) using the '--type' flag.`,
	Args: cobra.ExactArgs(1),
	Example: `
  Fetch all activities for a user:
    github-user-activity activity octocat

  Fetch only 'push' events for a user:
    github-user-activity activity octocat --type push`,
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]

		events, err := api.FetchUserActivity(username)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		filteredEvents := filters.FilterEventsByType(events, eventType)

		if len(filteredEvents) == 0 {
			fmt.Printf("No events found for user '%s' with type '%s'.\n", username, eventType)
			return
		}

		fmt.Printf("Recent activity for GitHub user '%s':\n", username)
		for _, event := range filteredEvents {
			fmt.Println(formatter.FormatEvent(event))
		}
	},
}

func init() {
	activityCmd.Flags().StringVar(&eventType, "type", "", "Filter events by type (e.g., push, issue, star)")
	rootCmd.AddCommand(activityCmd)
}
