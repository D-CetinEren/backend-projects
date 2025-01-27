// repository/repository.go
package repository

import "github.com/D-CetinEren/backend-projects/go/expense-tracker/internal/models"

type Repository interface {
	AddExpense(expense models.Expense) error
	GetExpenses() ([]models.Expense, error)
	UpdateExpense(id string, updatedExpense models.Expense) error // Changed to string ID
	DeleteExpense(id string) error                                // Changed to string ID
	// Remove GetNextID() and SaveLastID()
}
