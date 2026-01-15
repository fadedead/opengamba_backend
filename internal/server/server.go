package server

import (
	"net/http"
	"time"

	"github.com/fadedead/opengamba_backend/internal/auth"
	"github.com/fadedead/opengamba_backend/internal/rewards"
	"github.com/fadedead/opengamba_backend/internal/user"
	"github.com/fadedead/opengamba_backend/internal/websocket"
	"gorm.io/gorm"
)

type Config struct {
	Port          string
	DB            *gorm.DB
	UserHandler   *user.Handler
	AuthHandler   *auth.AuthHandler
	RewardHandler *rewards.Handler
	SocketHandler *websocket.Handler
}

type Server struct {
	port          string
	postgresDB    *gorm.DB
	userHandler   *user.Handler
	authHandler   *auth.AuthHandler
	rewardHandler *rewards.Handler
}

func NewServer(config Config) *http.Server {
	srv := &Server{
		port:          config.Port,
		postgresDB:    config.DB,
		userHandler:   config.UserHandler,
		authHandler:   config.AuthHandler,
		rewardHandler: config.RewardHandler,
	}

	return &http.Server{
		Addr:         ":" + config.Port,
		Handler:      srv.RegisterRoutes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
	}
}
