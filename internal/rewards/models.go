package rewards

import (
	"time"

	"github.com/fadedead/opengamba_backend/internal/user"
)

type Reward struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"index"`
	User      user.User `gorm:"foreignKey:UserID;references:ID"`
	Coins     int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
