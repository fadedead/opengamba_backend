package rewards

import "log/slog"

type Service interface {
	Save(reward *Reward) (Reward, error)
	GetById(id uint) (Reward, error)
	GetByUserId(userId uint) (Reward, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) Save(reward *Reward) (Reward, error) {
	savedReward, err := s.repo.Save(reward)
	if err != nil {
		slog.Error("Error saving reward", "reward", reward, "error", err)
		return savedReward, err
	}
	return savedReward, nil
}

func (s *service) GetById(id uint) (Reward, error) {
	reward, err := s.repo.GetById(id)
	if err != nil {
		slog.Error("Error getting reward with id", "id", id, "error", err)
		return reward, err
	}
	return reward, nil
}

func (s *service) GetByUserId(userId uint) (Reward, error) {
	reward, err := s.repo.GetByUserId(userId)
	if err != nil {
		slog.Error("Error getting reward with userId", "userId", userId, "error", err)
		return reward, err
	}
	return reward, nil
}
