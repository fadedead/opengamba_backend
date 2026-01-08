package user

type Service interface {
	SaveUser(user *User) (*User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) SaveUser(user *User) (*User, error) {
	err := s.repo.Save(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
