package rewards

import (
	"log/slog"

	"gorm.io/gorm"
)

type Repository interface {
	Save(reward *Reward) (Reward, error)
	Update(id uint, reward *Reward) (Reward, error)
	GetById(id uint) (Reward, error)
	GetByUserId(userId uint) (Reward, error)
	UpdateCoinsWithUserId(userId uint, coins int64) (Reward, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Save(reward *Reward) (Reward, error) {
	result := r.db.Create(reward)
	if result.Error != nil {
		slog.Error("Error creating reward", "error", result.Error)
		return *reward, result.Error
	}
	slog.Info("Created reward for user", "user", reward.UserID, "reward", reward.Coins)
	return *reward, nil
}

func (r *repository) Update(id uint, reward *Reward) (Reward, error) {
	foundReward, err := r.GetById(id)
	if err != nil {
		slog.Error("Error getting reward for update", "id", id, "error", err)
		return foundReward, err
	}

	foundReward.Coins = reward.Coins
	foundReward.UserID = reward.UserID

	result := r.db.Save(&foundReward)
	if result.Error != nil {
		slog.Error("Error updating the reward", "foundReward", foundReward, "reward", reward, "error", result.Error)
		return foundReward, result.Error
	}
	return foundReward, nil
}

func (r *repository) GetById(id uint) (Reward, error) {
	var foundReward Reward
	result := r.db.Where(&Reward{ID: id}).First(&foundReward)
	if result.Error != nil {
		slog.Error("Error finding reward with id", "id", id, "error", result.Error)
		return foundReward, result.Error
	}
	return foundReward, nil
}

func (r *repository) GetByUserId(userId uint) (Reward, error) {
	var foundReward Reward
	result := r.db.Preload("User").Where("user_id = ?", userId).First(&foundReward)
	if result.Error != nil {
		slog.Error("Error getting reward for user", "userId", userId, "error", result.Error)
		return foundReward, result.Error
	}
	return foundReward, nil
}

func (r *repository) UpdateCoinsWithUserId(userId uint, coins int64) (Reward, error) {
	foundReward, err := r.GetByUserId(userId)
	if err != nil {
		slog.Error("Error fetching the user for update", "userId", userId, "error", err)
		return foundReward, err
	}
	foundReward.Coins = coins
	result := r.db.Save(&foundReward)
	if result.Error != nil {
		slog.Error("Error updating coins for userId", "userId", userId, "coins", coins, "error", result.Error)
		return foundReward, result.Error
	}
	return foundReward, nil
}
