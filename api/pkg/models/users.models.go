package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string    `gorm:"size:255;not null"`
	Email     string    `gorm:"size:255;unique;not null"`
	Password  string    `gorm:"size:255;not null"`
	GoogleID  string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"not null"`
	Incomes   []Income  `gorm:"foreignKey:UserID;references:ID"`
	Expenses  []Expense `gorm:"foreignKey:UserID;references:ID"`
}
