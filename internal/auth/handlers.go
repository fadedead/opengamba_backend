package auth

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"

	"github.com/markbates/goth/gothic"
)

type Handler interface {
	Routes()
	Login()
	LoginCallback()
	Logout()
	GetUserSession()
}

type AuthHandler struct {
	service Service
}

func (h *AuthHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /auth/login/{provider}", h.Login)
	mux.HandleFunc("GET /auth/callback/{provider}", h.LoginCallback)
	mux.HandleFunc("GET /auth/logout/{provider}", h.Logout)
	mux.HandleFunc("GET /auth/me", h.GetUserSession)
}

func NewHandler(service Service) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /login/{provider}", h.Login)
	mux.HandleFunc("GET /callback/{provider}", h.LoginCallback)
	mux.HandleFunc("GET /logout/{provider}", h.Logout)
	mux.HandleFunc("GET /me", h.GetUserSession)
	return mux
}

func (h *AuthHandler) Login(res http.ResponseWriter, req *http.Request) {
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

func (h *AuthHandler) LoginCallback(res http.ResponseWriter, req *http.Request) {
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

	_, err = h.service.SaveUserToDatabase(user)
	if err != nil {
		slog.Error("Error saving user after callback", "error", err)
		http.Redirect(res, req, os.Getenv("FRONTEND_URL")+"/login?error=login_failed", http.StatusTemporaryRedirect)
	}
	slog.Info("User logged in", "user", user.Email)

	http.Redirect(res, req, os.Getenv("FRONTEND_URL")+"/home", http.StatusFound)
}

func (h *AuthHandler) Logout(res http.ResponseWriter, req *http.Request) {
	frontendUrl := os.Getenv("FRONTEND_URL")
	gothic.Logout(res, req)
	res.Header().Set("Location", frontendUrl)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func (h *AuthHandler) GetUserSession(res http.ResponseWriter, req *http.Request) {
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		slog.Error("Error getting user session", "error", err)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(res).Encode(map[string]string{"error": "Unauthorized"})
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(user)
}
