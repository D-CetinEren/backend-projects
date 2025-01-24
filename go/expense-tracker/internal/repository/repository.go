package repository

import "github.com/D-CetinEren/backend-projects/go/expense-tracker/internal/models"

// Repository defines methods to interact with the expense storage.
type Repository interface {
	AddExpense(expense models.Expense) error
	GetExpenses() ([]models.Expense, error)
	UpdateExpense(id int, updatedExpense models.Expense) error
	DeleteExpense(id int) error
}
