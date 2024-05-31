package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashedPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", fmt.Errorf("Can not generate hash from password %v", err)
	}

	return string(hash), nil
}

func CheckHashedPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
