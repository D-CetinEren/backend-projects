// models/expense.go
package models

import (
	"time"
)

type Expense struct {
	ID          string    `json:"id"` // Changed to UUID string
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
	Category    string    `json:"category"`
}
