package user

import (
	"log/slog"
)

type Service interface {
	Save(user *User) (User, error)
	GetById(id uint) (User, error)
}

type service struct {
	repo Repository
}

func NewService(repo *repository) *service {
	return &service{repo: repo}
}

func (s *service) Save(user *User) (User, error) {
	if user == nil || user.Email == "" {
		slog.Error("Error valid user struct must be passed with an email", "user", user, "error", ErrInvalidUserStruct)
		return User{}, ErrInvalidUserStruct
	}

	savedUser, err := s.repo.Save(user)
	if err != nil {
		slog.Error("Error saving user", "user", user, "error", err)
		return savedUser, err
	}
	slog.Info("Saved user", "user", savedUser)
	return savedUser, nil
}

func (s *service) GetById(id uint) (User, error) {
	if id == 0 {
		return User{}, ErrInvalidUserStruct
	}

	foundUser, err := s.repo.GetById(id)
	if err != nil {
		slog.Error("Error getting user with id", "id", id, "error", err)
		return User{}, err
	}
	return foundUser, nil
}
