package server

import (
	"net/http"
	"os"

	"github.com/fadedead/opengamba_backend/internal/auth"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// Auth routes
	mux.HandleFunc("GET /auth/{provider}", auth.Login)
	mux.HandleFunc("GET /auth/{provider}/callback", auth.LoginCallback)
	mux.HandleFunc("GET /logout/{provider}", auth.Logout)
	mux.HandleFunc("GET /auth/me", auth.GetUserSession)

	return s.enableCORS(mux)
}

func (s *Server) enableCORS(next http.Handler) http.Handler {
	frontendUrl := os.Getenv("FRONTEND_URL")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", frontendUrl)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
