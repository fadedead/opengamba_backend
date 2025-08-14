package main

import (
	"log/slog"
	"open_gamba/internal/database"
	"open_gamba/internal/router"
	"os"
	"time"

	ginslog "github.com/gin-contrib/slog"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	time.Local = time.UTC

	loadDotEnv()

	database.ConnectToSqlDb()
	if err := database.RunMigrations(); err != nil {
		slog.Error("Migration failed", "error", err)
		panic(err)
	}
}

func loadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		slog.Info("Error loading .env file, continuing without it")
	}
}

func main() {
	startApp()
}

func startApp() {
	// Set slog as the default logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// App using gin
	app := gin.New()
	app.Use(ginslog.SetLogger(
		ginslog.WithWriter(os.Stdout),
		ginslog.WithDefaultLevel(slog.LevelInfo),
		ginslog.WithMessage("HTTP Request"),
		ginslog.WithRequestHeader(true),
		ginslog.WithLogger(func(c *gin.Context, l *slog.Logger) *slog.Logger {
			return logger
		}),
	))
	app.Use(gin.Recovery())

	rootRouter := app.Group("/")
	router.AddRoutes(rootRouter)

	app.Run(":8080")
}
