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

func NewFileStore(filePath string) *FileStore {
	return &FileStore{filePath: filePath}
}

// AddExpense appends a new expense as a JSON line.
func (fs *FileStore) AddExpense(expense models.Expense) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	file, err := os.OpenFile(fs.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(expense)
}

// GetExpenses reads all JSON lines into a slice.
func (fs *FileStore) GetExpenses() ([]models.Expense, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	file, err := os.Open(fs.filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []models.Expense{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var expenses []models.Expense
	decoder := json.NewDecoder(file)
	for decoder.More() {
		var expense models.Expense
		if err := decoder.Decode(&expense); err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}
	return expenses, nil
}

// UpdateExpense and DeleteExpense require rewriting the entire file.
func (fs *FileStore) UpdateExpense(id string, updatedExpense models.Expense) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	expenses, err := fs.GetExpenses()
	if err != nil {
		return err
	}

	found := false
	for i, expense := range expenses {
		if expense.ID == id {
			expenses[i] = updatedExpense
			found = true
			break
		}
	}
	if !found {
		return errors.New("expense not found")
	}

	return fs.saveExpenses(expenses)
}

func (fs *FileStore) DeleteExpense(id string) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	expenses, err := fs.GetExpenses()
	if err != nil {
		return err
	}

	found := false
	for i, expense := range expenses {
		if expense.ID == id {
			expenses = append(expenses[:i], expenses[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		return errors.New("expense not found")
	}

	return fs.saveExpenses(expenses)
}

// saveExpenses writes all expenses as JSON lines.
func (fs *FileStore) saveExpenses(expenses []models.Expense) error {
	file, err := os.Create(fs.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	for _, expense := range expenses {
		if err := encoder.Encode(expense); err != nil {
			return err
		}
	}
	return nil
}
