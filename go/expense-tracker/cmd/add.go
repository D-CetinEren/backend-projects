package cmd

import (
	"fmt"
	"log"

	"github.com/D-CetinEren/backend-projects/go/expense-tracker/internal/repository"
	"github.com/D-CetinEren/backend-projects/go/expense-tracker/internal/services"

	"github.com/spf13/cobra"
)

var (
	description string
	amount      float64
	category    string
)

func init() {
	addCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the expense")
	addCmd.Flags().Float64VarP(&amount, "amount", "a", 0, "Amount of the expense")
	addCmd.Flags().StringVarP(&category, "category", "c", "Misc", "Category of the expense")
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new expense",
	Run: func(cmd *cobra.Command, args []string) {
		if description == "" || amount <= 0 {
			log.Println("Invalid input. Description and amount are required.")
			return
		}

		repo := repository.NewFileStore("data/expenses.json")
		service := services.NewExpenseService(repo)

		err := service.AddExpense(description, amount, category)
		if err != nil {
			log.Fatalf("Error adding expense: %v", err)
		}
		fmt.Println("Expense added successfully!")
	},
}
