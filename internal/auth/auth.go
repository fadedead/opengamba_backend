package auth

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

var (
	sessionSecret          = os.Getenv("SESSION_SECRET")
	googleAuthClientId     = os.Getenv("GOOGLE_AUTH_CLIENT_ID")
	googleAuthClientSecret = os.Getenv("GOOGLE_AUTH_CLIENT_SECRET")
	frontendUrl            = os.Getenv("FRONTEND_URL")
)

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

func SetupAuth() {
	store := sessions.NewCookieStore([]byte(sessionSecret))
	store.Options.HttpOnly = true
	store.Options.Path = "/"
	store.Options.Secure = false
	gothic.Store = store

	goth.UseProviders(
		google.New(googleAuthClientId, googleAuthClientSecret, frontendUrl+"/auth/callback/google"),
	)
}
