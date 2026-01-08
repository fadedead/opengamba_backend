package auth

import (
	"os"

	"github.com/fadedead/opengamba_backend/internal/user"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

type Service interface {
	SaveUserToDatabase(user goth.User) (*user.User, error)
}

type service struct {
	userService user.Service
}

func NewService(userService user.Service) *service {
	return &service{userService: userService}
}

func SetupAuth() {
	var (
		sessionSecret          = os.Getenv("SESSION_SECRET")
		googleAuthClientId     = os.Getenv("GOOGLE_AUTH_CLIENT_ID")
		googleAuthClientSecret = os.Getenv("GOOGLE_AUTH_CLIENT_SECRET")
		frontendUrl            = os.Getenv("FRONTEND_URL")
	)

	store := sessions.NewCookieStore([]byte(sessionSecret))
	store.Options.HttpOnly = true
	store.Options.Path = "/"
	store.Options.Secure = false
	gothic.Store = store
	goth.UseProviders(
		google.New(googleAuthClientId, googleAuthClientSecret, frontendUrl+"/auth/callback/google"),
	)
}

func (s *service) SaveUserToDatabase(gothUser goth.User) (*user.User, error) {
	user := &user.User{
		Email: gothUser.Email,
	}
	user, err := s.userService.SaveUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
