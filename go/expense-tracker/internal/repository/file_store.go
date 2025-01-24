package repository

import (
	"encoding/json"
	"errors"
	"os"
	"sync"

	"github.com/D-CetinEren/backend-projects/go/expense-tracker/internal/models"
)

type FileStore struct {
	filePath string
	mu       sync.Mutex
}

// NewFileStore creates a new FileStore instance.
func NewFileStore(filePath string) *FileStore {
	return &FileStore{filePath: filePath}
}

// AddExpense adds a new expense to the file.
func (fs *FileStore) AddExpense(expense models.Expense) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	expenses, err := fs.GetExpenses()
	if err != nil {
		return err
	}

	expenses = append(expenses, expense)
	return fs.saveExpenses(expenses)
}

// GetExpenses retrieves all expenses from the file.
func (fs *FileStore) GetExpenses() ([]models.Expense, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	file, err := os.Open(fs.filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []models.Expense{}, nil // Return empty if file doesn't exist
		}
		return nil, err
	}
	defer file.Close()

	var expenses []models.Expense
	if err := json.NewDecoder(file).Decode(&expenses); err != nil && err.Error() != "EOF" {
		return nil, err
	}
	return expenses, nil
}

// UpdateExpense updates an existing expense in the file.
func (fs *FileStore) UpdateExpense(id int, updatedExpense models.Expense) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	expenses, err := fs.GetExpenses()
	if err != nil {
		return err
	}

	for i, expense := range expenses {
		if expense.ID == id {
			expenses[i] = updatedExpense
			return fs.saveExpenses(expenses)
		}
	}
	return errors.New("expense not found")
}

// DeleteExpense deletes an expense by ID.
func (fs *FileStore) DeleteExpense(id int) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	expenses, err := fs.GetExpenses()
	if err != nil {
		return err
	}

	for i, expense := range expenses {
		if expense.ID == id {
			expenses = append(expenses[:i], expenses[i+1:]...)
			return fs.saveExpenses(expenses)
		}
	}
	return errors.New("expense not found")
}

// saveExpenses saves the expenses to the file.
func (fs *FileStore) saveExpenses(expenses []models.Expense) error {
	file, err := os.Create(fs.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(expenses)
}
