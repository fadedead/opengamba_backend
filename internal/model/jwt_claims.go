package model

import (
	"open_gamba/internal/user"

	"github.com/golang-jwt/jwt/v5"
)

type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type JwtCustomClaims struct {
	Username string            `json:"username"`
	Roles    []user.AccessRole `json:"roles"`
	jwt.RegisteredClaims
}
