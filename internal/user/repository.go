package user

import (
	"errors"
	"log/slog"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user *User) (User, error)
	GetById(id uint) (User, error)
	GetByEmail(email string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Save(user *User) (User, error) {
	foundUser, err := r.GetByEmail(user.Email)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error("Error confirming user exists", "userEmail", user.Email, "error", err)
		return foundUser, err
	}

	result := r.db.Create(user)
	if result.Error != nil {
		slog.Error("Error creating user", "error", result.Error)
		return foundUser, result.Error
	}
	slog.Info("Created user", "userId", user.ID, "userEmail", user.Email)
	return *user, result.Error
}

func (r *repository) GetById(id uint) (User, error) {
	var foundUser User
	result := r.db.Where(&User{ID: id}).First(&foundUser)
	if result.Error != nil {
		slog.Error("Error fetching user with id", "userId", id, "error", result.Error)
		return foundUser, result.Error
	}
	return foundUser, nil
}

func (r *repository) GetByEmail(email string) (User, error) {
	var foundUser User
	result := r.db.Where(&User{Email: email}).First(&foundUser)
	if result.Error != nil {
		slog.Error("Error fetching user with email", "email", email, "error", result.Error)
		return foundUser, result.Error
	}
	return foundUser, nil
}
