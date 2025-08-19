package model

type CoinFace string

const (
	Head CoinFace = "HEAD"
	Tail CoinFace = "TAIL"
)

type CoinFlipGame struct {
	BaseGame
	UserCoinFace    CoinFace `gorm:"type:varchar(10);not null"`
	WinningCoinFace CoinFace `gorm:"type:varchar(10);not null"`
}
