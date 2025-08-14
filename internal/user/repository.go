package user

import (
	"context"
	"log/slog"
	"open_gamba/internal/database"

	"gorm.io/gorm"
)

func SaveUser(user User) (*User, error) {
	ctx := context.Background()
	err := gorm.G[User](database.GetSqlDb()).Create(ctx, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetWithUsername(username string) (*User, error) {
	ctx := context.Background()
	fetchedUser, err := gorm.G[User](database.GetSqlDb()).Where(&User{Username: username}).First(ctx)
	if err != nil {
		return nil, err
	}
	slog.Info("User found with username", "username", username, "user", fetchedUser)
	return &fetchedUser, nil
}

func GetWithEmail(email string) (*User, error) {
	ctx := context.Background()
	fetchedUser, err := gorm.G[User](database.GetSqlDb()).Where(&User{Email: email}).First(ctx)
	if err != nil {
		return nil, err
	}
	slog.Info("User found with email", "email", email, "user", fetchedUser)
	return &fetchedUser, nil
}
