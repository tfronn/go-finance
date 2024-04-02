package models

import (
	"time"

	"github.com/google/uuid"
)

type Income struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Description string    `gorm:"not null"`
	Amount      float64   `gorm:"not null"`
	Category    string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
	UserID      uuid.UUID `gorm:"type:uuid;not null;index;foreignKey:UserID"`
}
