package cmd

import (
	"fmt"
	"log"

	"github.com/D-CetinEren/backend-projects/go/expense-tracker/internal/repository"
	"github.com/D-CetinEren/backend-projects/go/expense-tracker/internal/services"

	"github.com/spf13/cobra"
)

var month int

func init() {
	summaryCmd.Flags().IntVarP(&month, "month", "m", 0, "Summarize expenses for a specific month")
	rootCmd.AddCommand(summaryCmd)
}

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Show a summary of expenses",
	Run: func(cmd *cobra.Command, args []string) {
		repo := repository.NewFileStore("data/expenses.json")
		service := services.NewExpenseService(repo)

		if month == 0 {
			total, err := service.GetSummary()
			if err != nil {
				log.Fatalf("Error calculating summary: %v", err)
			}
			fmt.Printf("Total expenses: $%.2f\n", total)
		} else {
			// Handle monthly summary logic (to be implemented in service)
			fmt.Println("Monthly summaries will be added soon.")
		}
	},
}
