package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/D-CetinEren/backend-projects/go/Github-user-activity/internal/api"
	"github.com/D-CetinEren/backend-projects/go/Github-user-activity/internal/filters"
	"github.com/D-CetinEren/backend-projects/go/Github-user-activity/internal/formatter"
	"github.com/D-CetinEren/backend-projects/go/Github-user-activity/internal/models"
	"gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
)

var (
	cacheTTL   int    // Cache Time-To-Live in minutes
	maxPages   int    // Maximum number of pages to fetch from the GitHub API
	eventType  string // Filter events by type
	output     string // Output format: text, json, yaml
	outputFile string // File to save the output
)

var activityCmd = &cobra.Command{
	Use:   "activity [usernames...]",
	Short: "Fetch recent GitHub activity for one or more users",
	Long: `Fetch recent GitHub activity for the specified users and display it in the terminal or save it to a file.
This command supports caching to reduce API calls and various output formats.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		results := make(chan map[string]interface{}, len(args))
		ttl := time.Duration(cacheTTL) * time.Minute

		for _, username := range args {
			wg.Add(1)
			go func(username string) {
				defer wg.Done()

				events, err := api.FetchUserActivityWithCache(username, maxPages, ttl)
				if err != nil {
					results <- map[string]interface{}{
						"username": username,
						"error":    fmt.Sprintf("Error fetching activity: %v", err),
					}
					return
				}

				filteredEvents := filters.FilterEventsByType(events, eventType)
				results <- map[string]interface{}{
					"username": username,
					"events":   filteredEvents,
				}
			}(username)
		}

		wg.Wait()
		close(results)

		allResults := []map[string]interface{}{}
		for result := range results {
			allResults = append(allResults, result)
		}

		var outputData string
		switch output {
		case "json":
			outputData = formatJSON(allResults)
		case "yaml":
			outputData = formatYAML(allResults)
		default:
			outputData = formatText(allResults)
		}

		if outputFile != "" {
			if err := saveToFile(outputFile, outputData); err != nil {
				fmt.Printf("Error saving to file: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Output saved to file: %s\n", outputFile)
		} else {
			fmt.Println(outputData)
		}
	},
}

func formatJSON(data []map[string]interface{}) string {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("Error formatting JSON: %v\n", err)
		os.Exit(1)
	}
	return string(jsonData)
}

func formatYAML(data []map[string]interface{}) string {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		fmt.Printf("Error formatting YAML: %v\n", err)
		os.Exit(1)
	}
	return string(yamlData)
}

func formatText(data []map[string]interface{}) string {
	output := ""
	for _, result := range data {
		username := result["username"].(string)
		if err, exists := result["error"]; exists {
			output += fmt.Sprintf("Error fetching activity for user '%s': %s\n", username, err)
		} else {
			output += fmt.Sprintf("Recent activity for GitHub user '%s':\n", username)
			for _, event := range result["events"].([]interface{}) {
				// Assert the type to models.Event
				if eventMap, ok := event.(map[string]interface{}); ok {
					var event models.Event
					eventBytes, _ := json.Marshal(eventMap) // Re-marshal to JSON
					_ = json.Unmarshal(eventBytes, &event)  // Unmarshal back to models.Event
					output += formatter.FormatEvent(event) + "\n"
				} else {
					output += "Failed to parse event data.\n"
				}
			}
		}
	}
	return output
}

func saveToFile(filename, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	return err
}

func init() {
	// Define flags for the activity command
	activityCmd.Flags().IntVar(&cacheTTL, "cache-ttl", 10, "Cache time-to-live in minutes")
	activityCmd.Flags().IntVar(&maxPages, "max-pages", 1, "Maximum number of pages to fetch (default is 1)")
	activityCmd.Flags().StringVar(&eventType, "event-type", "", "Filter events by type (e.g., PushEvent, IssuesEvent)")
	activityCmd.Flags().StringVar(&output, "output", "text", "Output format: text, json, yaml")
	activityCmd.Flags().StringVar(&outputFile, "output-file", "", "File to save the output")

	// Add the activity command to the root command
	rootCmd.AddCommand(activityCmd)
}
