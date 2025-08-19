package model

import (
	"gorm.io/gorm"
)

type RewardConfig struct {
	gorm.Model
	BaseGameID     uint `gorm:"uniqueIndex"`
	WageAmount     int32
	RiskPercentage int
}
