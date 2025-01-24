package cmd

import (
	"fmt"
	"log"

	"github.com/D-CetinEren/backend-projects/go/expense-tracker/internal/repository"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all expenses",
	Run: func(cmd *cobra.Command, args []string) {
		repo := repository.NewFileStore("data/expenses.json")
		expenses, err := repo.GetExpenses()
		if err != nil {
			log.Fatalf("Error fetching expenses: %v", err)
		}

		if len(expenses) == 0 {
			fmt.Println("No expenses recorded.")
			return
		}

		fmt.Printf("%-5s %-15s %-20s %-10s\n", "ID", "Date", "Description", "Amount")
		for _, expense := range expenses {
			fmt.Printf("%-5d %-15s %-20s $%.2f\n", expense.ID, expense.Date.Format("2006-01-02"), expense.Description, expense.Amount)
		}
	},
}
