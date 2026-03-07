package utils

import (
	"crypto/rand"

)
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
func GenerateID() (string, error) {
	bytes := make([]byte, 8)
	random := make([]byte, 8)

	_, err := rand.Read(random)
	if err != nil {
		return "", err
	}

	for i := range bytes {
		bytes[i] = charset[random[i]%byte(len(charset))]
	}

	return string(bytes), nil
}