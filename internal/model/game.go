package model

import (
	"gorm.io/gorm"
)

type GameState string

const (
	Active   GameState = "ACTIVE"
	Inactive GameState = "INACTIVE"
)

type BaseGame struct {
	gorm.Model
	Username     string
	User         User          `gorm:"references:Username;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	RewardConfig *RewardConfig `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreditAmount int32
}
