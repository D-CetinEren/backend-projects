package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/D-CetinEren/backend-projects/go/Github-user-activity/internal/api"
	"github.com/D-CetinEren/backend-projects/go/Github-user-activity/internal/filters"
	"github.com/D-CetinEren/backend-projects/go/Github-user-activity/internal/formatter"
	"gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
)

var (
	cacheTTL     int    // Cache Time-To-Live in minutes
	maxPages     int    // Maximum number of pages to fetch from the GitHub API
	eventType    string // Filter events by type
	outputFormat string // Output format: text (default), json, or yaml
)

var activityCmd = &cobra.Command{
	Use:   "activity [username]",
	Short: "Fetch recent GitHub activity for a user",
	Long: `Fetch recent GitHub activity for the specified user and display it in the terminal.
This command supports caching to reduce API calls and allows output in different formats.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		ttl := time.Duration(cacheTTL) * time.Minute

		// Fetch user activity with caching
		events, err := api.FetchUserActivityWithCache(username, maxPages, ttl)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		// Filter events by type if specified
		filteredEvents := filters.FilterEventsByType(events, eventType)

		// Handle output format
		switch outputFormat {
		case "json":
			jsonData, err := json.MarshalIndent(filteredEvents, "", "  ")
			if err != nil {
				fmt.Printf("Error: failed to format output as JSON: %s\n", err)
				os.Exit(1)
			}
			fmt.Println(string(jsonData))
		case "yaml":
			yamlData, err := yaml.Marshal(filteredEvents)
			if err != nil {
				fmt.Printf("Error: failed to format output as YAML: %s\n", err)
				os.Exit(1)
			}
			fmt.Println(string(yamlData))
		default:
			// Default to text output
			if len(filteredEvents) == 0 {
				fmt.Printf("No events found for user '%s' with type '%s'.\n", username, eventType)
				return
			}

			fmt.Printf("Recent activity for GitHub user '%s':\n", username)
			for _, event := range filteredEvents {
				fmt.Println(formatter.FormatEvent(event))
			}
		}
	},
}

func init() {
	// Define flags for the activity command
	activityCmd.Flags().IntVar(&cacheTTL, "cache-ttl", 10, "Cache time-to-live in minutes")
	activityCmd.Flags().IntVar(&maxPages, "max-pages", 1, "Maximum number of pages to fetch (default is 1)")
	activityCmd.Flags().StringVar(&eventType, "event-type", "", "Filter events by type (e.g., PushEvent, IssuesEvent)")
	activityCmd.Flags().StringVar(&outputFormat, "output", "text", "Output format: text (default), json, or yaml")

	// Add the activity command to the root command
	rootCmd.AddCommand(activityCmd)
}
