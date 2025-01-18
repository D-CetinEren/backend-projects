package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/D-CetinEren/backend-projects/go/task-tracker/cmd"
	"github.com/D-CetinEren/backend-projects/go/task-tracker/internal/config"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Ensure required directories exist
	if err := os.MkdirAll(cfg.StoragePath, 0755); err != nil {
		log.Fatalf("Failed to create storage directory: %v", err)
	}
	if err := os.MkdirAll(cfg.LogPath, 0755); err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}

	// Set up signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		log.Println("Shutting down gracefully...")
		// Perform cleanup if needed
		os.Exit(0)
	}()

	// Execute the root command
	cmd.Execute()
}
