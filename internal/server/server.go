package server

import (
	"net/http"
	"time"
)

type Server struct {
	port string
}

func NewServer(port string) *http.Server {
	srv := &Server{port: port}

	return &http.Server{
		Addr:         ":" + port,
		Handler:      srv.RegisterRoutes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
	}
}
