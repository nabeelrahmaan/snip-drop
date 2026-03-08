package service

import (
	"codeDrop/internal/models"
	"codeDrop/internal/repository"
	"codeDrop/internal/utils"
	"errors"
	"time"
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
		CreatedAt: time.Now(),
	}

	return r.Repo.Create(&user)
}

func (r *AuthService) Login(email, pass string) (string, string, error) {

	user, err := r.Repo.FindByEmail(email)
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}

	if !utils.CheckPassword(user.Password, pass) {
		return "", "", errors.New("invalid credentials")
	}

	access, err := utils.GenerateAccess(user)
	if err != nil {
		return "", "", errors.New("failed to create access token")
	}

	refresh, err := utils.GenerateRefresh(user)
	if err != nil {
		return "", "", errors.New("failed to create refresh token")
	}

	ref := models.RefreshToken{
		UserID: user.ID,
		Token: refresh,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	err = r.Repo.CreateRefresh(&ref)
	if err != nil {
		return "", "", errors.New("failed to store refresh token")
	}

	return access, refresh, nil
}