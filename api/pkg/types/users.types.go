package types

import (
	"time"

	"github.com/google/uuid"
)

type UserDTO struct {
	ID        uuid.UUID   `json:"id"`
	Name      string      `json:"name"`
	Email     string      `json:"email"`
	Password  string      `json:"password,omitempty"`
	GoogleID  string      `json:"google_id,omitempty"`
	CreatedAt time.Time   `json:"created_at"`
	Incomes   *IncomeDTO  `json:"incomes,omitempty"`
	Expenses  *ExpenseDTO `json:"expenses,omitempty"`
}
