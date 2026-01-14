package server

import (
	"net/http"
	"os"
)

// RouteRegistrar interface to be implemented by each service package
type RouteRegistrar interface {
	RegisterRoutes(mux *http.ServeMux)
}

func (s *Server) RegisterRoutes() http.Handler {
	mainMux := http.NewServeMux()

	// Use StripPrefix approach to avoid route conflicts
	mainMux.Handle("/user/", http.StripPrefix("/user", s.userHandler.Routes()))
	mainMux.Handle("/auth/", http.StripPrefix("/auth", s.authHandler.Routes()))
	mainMux.Handle("/rewards/", http.StripPrefix("/rewards", s.rewardHandler.Routes()))

	return s.enableCORS(mainMux)
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
