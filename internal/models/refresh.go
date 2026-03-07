package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshToken struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	UserID uuid.UUID `gorm:"type:uuid;index"`
	User   User      `gorm:"foreignKey:UserID"`

	Token     string `gorm:"uniqueIndex"`
	ExpiresAt time.Time
	CreatedAt time.Time
}

func (r *RefreshToken) GenerateID(tx *gorm.DB) error {
	r.ID = uuid.New()
	return nil
}