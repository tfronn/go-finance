package types

import (
	"time"

	"github.com/google/uuid"
)

type IncomeDTO struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Amount      float64   `json:"amount,omitempty"`
	Description string    `json:"description,omitempty"`
	Category    string    `json:"category,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UserID      uuid.UUID `json:"user_id,omitempty"`
}
