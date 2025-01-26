package services

import (
	"errors"
	"time"

	"github.com/D-CetinEren/backend-projects/go/expense-tracker/internal/models"
	"github.com/D-CetinEren/backend-projects/go/expense-tracker/internal/repository"
)

type ExpenseService struct {
	repo repository.Repository
}

// NewExpenseService creates a new ExpenseService instance.
func NewExpenseService(repo repository.Repository) *ExpenseService {
	return &ExpenseService{repo: repo}
}

// AddExpense adds a new expense with the provided details.
func (es *ExpenseService) AddExpense(description string, amount float64, category string) error {
	if description == "" || amount <= 0 {
		return errors.New("invalid description or amount")
	}

	nextID, err := es.repo.GetNextID()
	if err != nil {
		return err
	}

	newExpense := models.Expense{
		ID:          nextID,
		Description: description,
		Amount:      amount,
		Date:        time.Now(),
		Category:    category,
	}

	if err := es.repo.AddExpense(newExpense); err != nil {
		return err
	}

	return es.repo.SaveLastID(nextID)
}

// GetSummary calculates the total of all expenses.
func (es *ExpenseService) GetSummary() (float64, error) {
	expenses, err := es.repo.GetExpenses()
	if err != nil {
		return 0, err
	}

	total := 0.0
	for _, expense := range expenses {
		total += expense.Amount
	}
	return total, nil
}
