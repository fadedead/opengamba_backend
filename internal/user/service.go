package user

import (
	"errors"
	"log/slog"

	"gorm.io/gorm"
)

func IsUserPresent(username string, email string) (bool, error) {
	foundUserWithUsername, err := GetWithUsername(username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error("Error fetching user with username", "username", err, "error", err)
		return false, err
	}
	if foundUserWithUsername != nil {
		return true, nil
	}

	foundUserWithEmail, err := GetWithEmail(email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error("Error fetching user with email", "email", err, "error", err)
		return false, err
	}
	return foundUserWithEmail != nil, nil
}
