package user

import (
	"log/slog"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user *User) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Save(user *User) error {
	err := r.db.Create(user).Error
	if err != nil {
		slog.Error("Error creating user", "error", err)
		return err
	}
	slog.Info("Created user", "userId", user.ID, "userEmail", user.Email)
	return nil
}
