package auth

import (
	"time"
)

type SignUpRequest struct {
	Email    string    `json:"email" binding:"required,email"`
	Password string    `json:"password" binding:"required,min=8,max=16"`
	Username string    `json:"username" binding:"required,alphanum,min=1,max=16"`
	Birthday time.Time `json:"birthday"`
}

type LoginRequest struct {
	Password string `json:"password" binding:"required,min=8,max=16"`
	Username string `json:"username" binding:"required,alphanum,min=1,max=16"`
}
