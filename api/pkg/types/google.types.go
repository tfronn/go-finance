package types

import (
	"gofinance/api/pkg/models"
	"time"

	"github.com/google/uuid"
)

type GoogleUserInfo struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Verified   bool   `json:"verified_email"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
	Locale     string `json:"locale"`
}

type GoogleUserDTO struct {
	ID        uuid.UUID       `json:"id"`
	Name      string          `json:"name"`
	Email     string          `json:"email"`
	Password  string          `json:"password,omitempty"`
	GoogleID  string          `json:"google_id,omitempty"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Incomes   interface{}     `json:"incomes,omitempty"`
	Expenses  interface{}     `json:"expenses,omitempty"`
	UserID    uuid.UUID       `json:"-"`
	IncomeID  uuid.UUID       `json:"-"`
	ExpenseID uuid.UUID       `json:"-"`
	Income    *models.Income  `json:"-"`
	Expense   *models.Expense `json:"-"`
}
