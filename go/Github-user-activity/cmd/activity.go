package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
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
	outputFile   string // File to write the output
)

var activityCmd = &cobra.Command{
	Use:   "activity [username]",
	Short: "Fetch recent GitHub activity for a user",
	Long: `Fetch recent GitHub activity for the specified user and display it in the terminal or write it to a file.
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

		// Generate output
		var outputData string
		switch outputFormat {
		case "json":
			jsonData, err := json.MarshalIndent(filteredEvents, "", "  ")
			if err != nil {
				fmt.Printf("Error: failed to format output as JSON: %s\n", err)
				os.Exit(1)
			}
			outputData = string(jsonData)
		case "yaml":
			yamlData, err := yaml.Marshal(filteredEvents)
			if err != nil {
				fmt.Printf("Error: failed to format output as YAML: %s\n", err)
				os.Exit(1)
			}
			outputData = string(yamlData)
		default:
			// Default to text output
			if len(filteredEvents) == 0 {
				fmt.Printf("No events found for user '%s' with type '%s'.\n", username, eventType)
				return
			}

			var formattedEvents []string
			for _, event := range filteredEvents {
				formattedEvents = append(formattedEvents, formatter.FormatEvent(event))
			}
			outputData = fmt.Sprintf("Recent activity for GitHub user '%s':\n%s", username,
				strings.Join(formattedEvents, "\n"))
		}

		// Write output to file or display in terminal
		if outputFile != "" {
			if err := os.WriteFile(outputFile, []byte(outputData), 0644); err != nil {
				fmt.Printf("Error: failed to write output to file '%s': %s\n", outputFile, err)
				os.Exit(1)
			}
			fmt.Printf("Output written to file '%s'\n", outputFile)
		} else {
			fmt.Println(outputData)
		}
	},
}

func init() {
	// Define flags for the activity command
	activityCmd.Flags().IntVar(&cacheTTL, "cache-ttl", 10, "Cache time-to-live in minutes")
	activityCmd.Flags().IntVar(&maxPages, "max-pages", 1, "Maximum number of pages to fetch (default is 1)")
	activityCmd.Flags().StringVar(&eventType, "event-type", "", "Filter events by type (e.g., PushEvent, IssuesEvent)")
	activityCmd.Flags().StringVar(&outputFormat, "output", "text", "Output format: text (default), json, or yaml")
	activityCmd.Flags().StringVar(&outputFile, "output-file", "", "File to write the output")

	// Add the activity command to the root command
	rootCmd.AddCommand(activityCmd)
}
