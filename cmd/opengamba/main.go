package main

import (
	"log/slog"
	"os"

	"github.com/fadedead/opengamba_backend/internal/auth"
	"github.com/fadedead/opengamba_backend/internal/database"
	"github.com/fadedead/opengamba_backend/internal/rewards"
	"github.com/fadedead/opengamba_backend/internal/server"
	"github.com/fadedead/opengamba_backend/internal/user"
	"github.com/joho/godotenv"
	"github.com/pressly/goose"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading dotenv", "error", err)
		os.Exit(1)
	}

	postgresDB, err := database.ConnectToPostgresDB()
	if err != nil {
		slog.Error("Error connecting to Postgres", "error", err)
		os.Exit(1)
	}
	database.NewPostgresRepository(postgresDB)

	sqlDB, err := postgresDB.DB()
	if err != nil {
		slog.Error("Error getting sql.DB from GORM", "error", err)
		os.Exit(1)
	}
	if err := goose.SetDialect("postgres"); err != nil {
		slog.Error("Error setting up goose", "error", err)
	}
	if err := goose.Up(sqlDB, "migrations"); err != nil {
		slog.Error("Error performing the migrations", "error", err)
	}
	slog.Info("Migrations complete. Starting app...")

	// Initialize user service dependencies
	userRepository := user.NewRepository(postgresDB)
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	// Initialize auth service dependencies
	auth.SetupAuth()
	authService := auth.NewService(userService)
	authHandler := auth.NewHandler(authService)

	// Initialize rewards service dependencies
	rewardRepository := rewards.NewRepository(postgresDB)
	rewardService := rewards.NewService(rewardRepository)
	rewardHandler := rewards.NewHandler(rewardService)

	// Create server with all dependencies injected
	config := server.Config{
		Port:          "8080",
		DB:            postgresDB,
		UserHandler:   userHandler,
		AuthHandler:   authHandler,
		RewardHandler: rewardHandler,
	}
	srv := server.NewServer(config)
	slog.Info("Server starting...")
	err = srv.ListenAndServe()
	if err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
