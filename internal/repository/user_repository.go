package repository

import (
	"codeDrop/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user *models.User

	err := r.DB.Where("email=?", email).First(&user).Error
	
	return user, err
}