package utils

import (
	"codeDrop/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var AccessSecret = []byte("access_secret")
var RefreshSecret = []byte("refresh_secret")

func GenerateAccess(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role": user.Role,
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(AccessSecret)
}

func GenerateRefresh (user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role": user.Role,
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(RefreshSecret)
}