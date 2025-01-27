package cmd

import (
	"fmt"
	"log"

	"github.com/D-CetinEren/backend-projects/go/expense-tracker/internal/repository"

	"github.com/spf13/cobra"
)

var id string

func init() {
	deleteCmd.Flags().StringVarP(&id, "id", "i", "", "ID of the expense to delete")
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an expense by ID",
	Run: func(cmd *cobra.Command, args []string) {
		if id == "" {
			log.Println("Invalid ID. Please provide a valid expense ID.")
			return
		}

		repo := repository.NewFileStore("data/expenses.json")
		err := repo.DeleteExpense(id)
		if err != nil {
			log.Fatalf("Error deleting expense: %v", err)
		}

		fmt.Println("Expense deleted successfully!")
	},
}
