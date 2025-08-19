package coinflip

import (
	"hash/fnv"
	"math/rand"
	"open_gamba/internal/model"
	"os"
	"time"
)

func getFlipResult() model.CoinFace {
	h := fnv.New64a()
	h.Write([]byte(os.Getenv("COIN_FLIP_GAME_RANDOM_SEED")))
	stringHash := int64(h.Sum64())

	r := rand.New(rand.NewSource(time.Now().UnixNano() ^ stringHash))
	randomNum := r.Intn(2)

	if randomNum == 0 {
		return model.Head
	}
	return model.Tail
}
