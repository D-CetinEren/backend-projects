package models

import (
	"testing"
	"time"
)

func TestExpense(t *testing.T) {
	now := time.Now()
	expense := Expense{
		ID:          "test-id",
		Description: "Test Expense",
		Amount:      100.0,
		Date:        now,
		Category:    "Test",
	}

	t.Run("Check Expense Fields", func(t *testing.T) {
		if expense.ID != "test-id" {
			t.Errorf("Expected ID 'test-id', got %s", expense.ID)
		}
		if expense.Description != "Test Expense" {
			t.Errorf("Expected Description 'Test Expense', got %s", expense.Description)
		}
		if expense.Amount != 100.0 {
			t.Errorf("Expected Amount 100.0, got %f", expense.Amount)
		}
		if !expense.Date.Equal(now) {
			t.Errorf("Expected Date %v, got %v", now, expense.Date)
		}
		if expense.Category != "Test" {
			t.Errorf("Expected Category 'Test', got %s", expense.Category)
		}
	})
}
