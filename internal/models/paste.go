package models

import (
	"time"

	"github.com/google/uuid"
)

type Paste struct {
	ID         string `gorm:"primaryKey"`

	UserId     uuid.UUID `gorm:"type:uuid;index"`
	User User `gorm:"foreignkey:UserID"`

	ObjectKey  string `gorm:"not null"`
	Visibility string
	CreatedAt  time.Time
	ExpiresAt  *time.Time
}

