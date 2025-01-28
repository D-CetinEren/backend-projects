package services

import (
	"testing"
	"time"

	"github.com/D-CetinEren/backend-projects/go/expense-tracker/internal/models"
)

// MockRepository implements repository.Repository interface for testing
type MockRepository struct {
	expenses []models.Expense
	err      error
}

func (m *MockRepository) AddExpense(expense models.Expense) error {
	if m.err != nil {
		return m.err
	}
	m.expenses = append(m.expenses, expense)
	return nil
}

func (m *MockRepository) GetExpenses() ([]models.Expense, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.expenses, nil
}

func (m *MockRepository) UpdateExpense(id string, expense models.Expense) error {
	if m.err != nil {
		return m.err
	}
	for i, e := range m.expenses {
		if e.ID == id {
			m.expenses[i] = expense
			return nil
		}
	}
	return nil
}

func (m *MockRepository) DeleteExpense(id string) error {
	if m.err != nil {
		return m.err
	}
	for i, e := range m.expenses {
		if e.ID == id {
			m.expenses = append(m.expenses[:i], m.expenses[i+1:]...)
			return nil
		}
	}
	return nil
}

func TestAddExpense(t *testing.T) {
	tests := []struct {
		name        string
		description string
		amount      float64
		category    string
		wantErr     bool
	}{
		{
			name:        "Valid expense",
			description: "Groceries",
			amount:      50.0,
			category:    "Food",
			wantErr:     false,
		},
		{
			name:        "Empty description",
			description: "",
			amount:      50.0,
			category:    "Food",
			wantErr:     true,
		},
		{
			name:        "Zero amount",
			description: "Groceries",
			amount:      0,
			category:    "Food",
			wantErr:     true,
		},
		{
			name:        "Negative amount",
			description: "Groceries",
			amount:      -50.0,
			category:    "Food",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockRepository{}
			service := NewExpenseService(mockRepo)

			err := service.AddExpense(tt.description, tt.amount, tt.category)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddExpense() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if len(mockRepo.expenses) != 1 {
					t.Errorf("AddExpense() didn't store expense, got %d expenses", len(mockRepo.expenses))
				}
			}
		})
	}
}

func TestGetSummary(t *testing.T) {
	tests := []struct {
		name      string
		expenses  []models.Expense
		wantTotal float64
		wantErr   bool
	}{
		{
			name: "Multiple expenses",
			expenses: []models.Expense{
				{
					ID:          "1",
					Description: "Expense 1",
					Amount:      100.0,
					Date:        time.Now(),
					Category:    "Food",
				},
				{
					ID:          "2",
					Description: "Expense 2",
					Amount:      50.0,
					Date:        time.Now(),
					Category:    "Transport",
				},
			},
			wantTotal: 150.0,
			wantErr:   false,
		},
		{
			name:      "No expenses",
			expenses:  []models.Expense{},
			wantTotal: 0,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockRepository{expenses: tt.expenses}
			service := NewExpenseService(mockRepo)

			total, err := service.GetSummary()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSummary() error = %v, wantErr %v", err, tt.wantErr)
			}

			if total != tt.wantTotal {
				t.Errorf("GetSummary() = %v, want %v", total, tt.wantTotal)
			}
		})
	}
}
