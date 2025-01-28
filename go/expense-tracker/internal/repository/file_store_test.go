package repository

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/D-CetinEren/backend-projects/go/expense-tracker/internal/models"
)

func TestFileStore(t *testing.T) {
	// Create temporary file for testing
	tempDir := t.TempDir()
	tempFile := filepath.Join(tempDir, "test_expenses.json")

	store := NewFileStore(tempFile)

	// Test expense
	testExpense := models.Expense{
		ID:          "test-id",
		Description: "Test expense",
		Amount:      100.0,
		Date:        time.Now(),
		Category:    "Test",
	}

	// Test AddExpense
	t.Run("AddExpense", func(t *testing.T) {
		err := store.AddExpense(testExpense)
		if err != nil {
			t.Errorf("AddExpense() error = %v", err)
		}

		// Verify file exists
		if _, err := os.Stat(tempFile); os.IsNotExist(err) {
			t.Error("AddExpense() didn't create file")
		}
	})

	// Test GetExpenses
	t.Run("GetExpenses", func(t *testing.T) {
		expenses, err := store.GetExpenses()
		if err != nil {
			t.Errorf("GetExpenses() error = %v", err)
		}

		if len(expenses) != 1 {
			t.Errorf("GetExpenses() got %d expenses, want 1", len(expenses))
		}

		if expenses[0].ID != testExpense.ID {
			t.Errorf("GetExpenses() got ID = %v, want %v", expenses[0].ID, testExpense.ID)
		}
	})

	// Test DeleteExpense
	t.Run("DeleteExpense", func(t *testing.T) {
		err := store.DeleteExpense(testExpense.ID)
		if err != nil {
			t.Errorf("DeleteExpense() error = %v", err)
		}

		expenses, _ := store.GetExpenses()
		if len(expenses) != 0 {
			t.Errorf("DeleteExpense() didn't remove expense, got %d expenses", len(expenses))
		}
	})
}
