package types

import (
	"time"

	"github.com/google/uuid"
)

type ExpenseDTO struct {
	ID        uuid.UUID `json:"id"`
	Amount    float64   `json:"amount"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}
