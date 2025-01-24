package cmd

import (
	"fmt"
	"log"

	"github.com/D-CetinEren/backend-projects/go/expense-tracker/internal/repository"

	"github.com/spf13/cobra"
)

var id int

func init() {
	deleteCmd.Flags().IntVarP(&id, "id", "i", 0, "ID of the expense to delete")
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an expense by ID",
	Run: func(cmd *cobra.Command, args []string) {
		if id <= 0 {
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
