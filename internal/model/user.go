package model

import "time"

type AccessRole string

const (
	UserRole  AccessRole = "USER"
	AdminRole AccessRole = "ADMIN"
)

type User struct {
	ID          uint
	Username    string `gorm:"uniqueIndex"`
	Email       string `gorm:"uniqueIndex"`
	Birthday    *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Password    string
	AccessRoles []AccessRole `gorm:"type:text[]"`
}
