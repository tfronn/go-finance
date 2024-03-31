package entities

import (
	"time"

	"github.com/google/uuid"
)

type Income struct {
	ID        uuid.UUID
	Amount    float64
	Category  string
	CreatedAt time.Time
}

func NewIncome(amount float64, category string) Income {
	return Income{
		ID:        uuid.New(),
		Amount:    amount,
		Category:  category,
		CreatedAt: time.Now(),
	}
}
