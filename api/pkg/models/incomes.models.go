package models

import (
	"time"

	"github.com/google/uuid"
)

type Income struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Amount    float64   `gorm:"not null"`
	Category  string
	CreatedAt time.Time
	UserID    uuid.UUID `gorm:"type:uuid;not null;index"`
}
