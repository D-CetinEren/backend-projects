package models

import "time"

// Expense represents a single financial expense.
type Expense struct {
	ID          int       `json:"id"`          // Unique identifier for the expense
	Description string    `json:"description"` // Description of the expense
	Amount      float64   `json:"amount"`      // Amount spent
	Date        time.Time `json:"date"`        // Date of the expense
	Category    string    `json:"category"`    // Optional: Category of the expense
}
