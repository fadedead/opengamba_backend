package util

import (
	"fmt"
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func IsPasswordMatch(storedPassword string, enteredPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(enteredPassword))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, fmt.Errorf("Invalid password")
		}
		slog.Error("Error comparing password", "error", err)
		return false, err
	}
	return true, nil
}
