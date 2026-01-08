package user

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
