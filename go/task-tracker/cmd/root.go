package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task-tracker",
	Short: "A simple task tracking CLI application",
	Long:  `Manage and track your tasks through the command line interface.`,
}

// Execute initializes the root command and sets up logging.
func Execute() {
	setupLogging() // Initialize logging system
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing root command: %v", err)
	}
}

// setupLogging configures the logging system.
func setupLogging() {
	logFile, err := os.OpenFile("logs/task-tracker.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error setting up log file: %v\n", err)
		os.Exit(1)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Task tracker started.")
}

func init() {
	// Global flags and subcommands can be initialized here
}
