package server

import (
	"net/http"
	"time"

	"gorm.io/gorm"
)

type Server struct {
	port       string
	postgresDB *gorm.DB
}

func NewServer(port string, postgresDB *gorm.DB) *http.Server {
	srv := &Server{port: port, postgresDB: postgresDB}

	return &http.Server{
		Addr:         ":" + port,
		Handler:      srv.RegisterRoutes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
	}
}
