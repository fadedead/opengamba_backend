package auth

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"

	"github.com/markbates/goth/gothic"
)

func Login(res http.ResponseWriter, req *http.Request) {
	provider := req.PathValue("provider")
	q := req.URL.Query()
	q.Add("provider", provider)
	req.URL.RawQuery = q.Encode()

	if _, err := gothic.CompleteUserAuth(res, req); err == nil {
		http.Redirect(res, req, os.Getenv("FRONTEND_URL")+"/home", http.StatusFound)
	} else {
		gothic.BeginAuthHandler(res, req)
	}
}

func LoginCallback(res http.ResponseWriter, req *http.Request) {
	provider := req.PathValue("provider")
	q := req.URL.Query()
	q.Add("provider", provider)
	req.URL.RawQuery = q.Encode()

	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		slog.Error("Error in login callback", "error", err)
		http.Redirect(res, req, os.Getenv("FRONTEND_URL")+"/login?error=auth_failed", http.StatusTemporaryRedirect)
		return
	}

	slog.Info("User logged in", "user", user.Email)
	// TODO: save user to DB

	http.Redirect(res, req, os.Getenv("FRONTEND_URL")+"/home", http.StatusFound)
}

func Logout(res http.ResponseWriter, req *http.Request) {
	frontendUrl := os.Getenv("FRONTEND_URL")
	gothic.Logout(res, req)
	res.Header().Set("Location", frontendUrl)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func GetUserSession(res http.ResponseWriter, req *http.Request) {
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(res).Encode(map[string]string{"error": "Unauthorized"})
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(user)
}
