package server

import (
	"net/http"
	"os"

	"github.com/fadedead/opengamba_backend/internal/auth"
	"github.com/fadedead/opengamba_backend/internal/user"
	"gorm.io/gorm"
)

func (s *Server) RegisterRoutes() http.Handler {
	mainMux := http.NewServeMux()

	// Add your handler constructor here
	userService := setupUserHandler(s.postgresDB, mainMux)
	setupAuthHandler(userService, mainMux)

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

func setupAuthHandler(userService user.Service, mainMux *http.ServeMux) auth.Service {
	auth.SetupAuth()
	authService := auth.NewService(userService)
	authHandler := auth.NewHandler(authService)
	mainMux.Handle("/auth/", http.StripPrefix("/auth", authHandler.Routes()))
	return authService
}

func setupUserHandler(db *gorm.DB, mainMux *http.ServeMux) user.Service {
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	mainMux.Handle("/user/", http.StripPrefix("/user", userHandler.Routes()))
	return userService
}
