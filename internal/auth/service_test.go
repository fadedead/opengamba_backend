package auth

import (
	"open_gamba/internal/user"
	"testing"
)

func TestGenerateTokenWithValidUser(t *testing.T) {
	t.Setenv("AUTH_JWT_SECRET_KEY", "dummytoken1234")
	newUser := user.User{
		ID:          1,
		Email:       "test@gmail.com",
		Password:    "asdasdasd",
		Username:    "testuser",
		Birthday:    nil,
		AccessRoles: []user.AccessRole{user.UserRole},
	}

	token, err := GenerateToken(&newUser)
	if err != nil {
		t.Error("Valid user should generate token", "user", newUser, "error", err)
	}

	if token == "" {
		t.Error("Empty token generated for user", "user", newUser, "token", token)
	}
}

func TestDontGenerateTokenWithInvalidUsername(t *testing.T) {
	t.Setenv("AUTH_JWT_SECRET_KEY", "dummytoken1234")
	newUser := user.User{
		ID:          1,
		Email:       "test@gmail.com",
		Password:    "asdasdasd",
		Username:    "",
		Birthday:    nil,
		AccessRoles: []user.AccessRole{user.UserRole},
	}

	_, err := GenerateToken(&newUser)
	if err == nil {
		t.Error("No error for invalid user", "user", newUser)
	}
}

func TestDontGenerateTokenWithInvalidRoles(t *testing.T) {
	t.Setenv("AUTH_JWT_SECRET_KEY", "dummytoken1234")
	newUser := user.User{
		ID:          1,
		Email:       "test@gmail.com",
		Password:    "asdasdasd",
		Username:    "testuser",
		Birthday:    nil,
		AccessRoles: []user.AccessRole{},
	}

	_, err := GenerateToken(&newUser)
	if err == nil {
		t.Error("No error for invalid roles", "user", newUser)
	}
}

func TestValidTokenThrowsNoError(t *testing.T) {
	t.Setenv("AUTH_JWT_SECRET_KEY", "dummytoken1234")
	newUser := user.User{
		ID:          1,
		Email:       "test@gmail.com",
		Password:    "asdasdasd",
		Username:    "testuser",
		Birthday:    nil,
		AccessRoles: []user.AccessRole{user.UserRole},
	}

	token, _ := GenerateToken(&newUser)
	_, err := ValidateToken(token)
	if err != nil {
		t.Error("Valid token should not throw an error", "user", newUser, "token", token, "error", err)
	}
}

func TestInvalidTokenThrowsError(t *testing.T) {
	t.Setenv("AUTH_JWT_SECRET_KEY", "dummytoken1234")
	newUser := user.User{
		ID:          1,
		Email:       "test@gmail.com",
		Password:    "asdasdasd",
		Username:    "testuser",
		Birthday:    nil,
		AccessRoles: []user.AccessRole{user.UserRole},
	}

	token, _ := GenerateToken(&newUser)
	_, err := ValidateToken("RandomString")
	if err == nil {
		t.Error("Invalid token should throw an error", "user", newUser, "token", token, "error", err)
	}
}
