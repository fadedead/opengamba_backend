package main

import (
	"log/slog"
	"os"

	"github.com/fadedead/opengamba_backend/internal/auth"
	"github.com/fadedead/opengamba_backend/internal/database"
	"github.com/fadedead/opengamba_backend/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading dotenv", "error", err)
		os.Exit(1)
	}

	err = database.ConnectToPostgresDB()
	if err != nil {
		slog.Error("Error connecting to Postgres", "error", err)
		os.Exit(1)
	}

	auth.SetupAuth()

	srv := server.NewServer("8080")
	slog.Info("Server starting...")
	err = srv.ListenAndServe()
	if err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
