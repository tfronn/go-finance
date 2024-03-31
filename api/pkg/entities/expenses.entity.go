package entities

import (
	"time"

	"github.com/google/uuid"
)

type Expense struct {
	ID        uuid.UUID
	Amount    float64
	Category  string
	CreatedAt time.Time
}

func NewExpense(amount float64, category string) Expense {
	return Expense{
		ID:        uuid.New(),
		Amount:    amount,
		Category:  category,
		CreatedAt: time.Now(),
	}
}
