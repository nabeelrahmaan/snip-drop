package service

import (
	"codeDrop/internal/models"
	"codeDrop/internal/repository"
	"codeDrop/internal/utils"
	"errors"
)

type AuthService struct {
	Repo *repository.UserRepository
}

func (r *AuthService) Signup(name, email, password string) error {
	_, err := r.Repo.FindByEmail(email)
	if err == nil {
		return errors.New("User already exist")
	}

	hash, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := models.User{
		Username: name,
		Email: email,
		Password: hash,
	}

	return r.Repo.Create(&user)
}