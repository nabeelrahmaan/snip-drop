package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"unique;not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"default:user"`
	CreatedAt time.Time
}

func (u *User) GenerateID(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
