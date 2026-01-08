package main

import (
	"log/slog"
	"os"

	"github.com/fadedead/opengamba_backend/internal/database"
	"github.com/fadedead/opengamba_backend/internal/server"
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
	database.NewPostgresRegistry(postgresDB)

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

	srv := server.NewServer("8080", postgresDB)
	slog.Info("Server starting...")
	err = srv.ListenAndServe()
	if err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
