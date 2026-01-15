package websocket

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"sync"

	"github.com/fadedead/opengamba_backend/internal/user"
	"github.com/gorilla/websocket"
)

type Service interface {
	registerClientConnection(userId string, conn *websocket.Conn)
	SendAll(msg message)
	SendWithClientId(clientId string, msg message) error
}

type service struct {
	userService user.Service
	server      *server
}

type message struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type client struct {
	id   string
	conn *websocket.Conn
	send chan message
}

type server struct {
	clients    map[string]*client
	register   chan *client
	unregister chan *client
	mu         sync.Mutex
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func NewService(userService user.Service) *service {
	return &service{
		userService: userService,
		server:      newServer(),
	}
}

func newServer() *server {
	return &server{
		clients:    make(map[string]*client),
		register:   make(chan *client),
		unregister: make(chan *client),
	}
}

func (s *server) eventHandler(client *client, msg message) {
	slog.Info("Received message from client", "client", client.id, "message", msg)

	switch msg.Type {
	default:
		slog.Error("Unknown even type")
	}
}

func (c *client) readPump(s *server) {
	defer func() {
		s.unregister <- c
		c.conn.Close()
	}()

	for {
		_, raw, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		var msg message
		if err := json.Unmarshal(raw, &msg); err != nil {
			slog.Error("Json decoding error", err)
			continue
		}

		s.eventHandler(c, msg)
	}
}

func (c *client) writePump() {
	for msg := range c.send {
		c.conn.WriteJSON(msg)
	}
}

func (s *service) registerClientConnection(userId string, conn *websocket.Conn) {
	client := &client{
		id: userId, conn: conn, send: make(chan message),
	}
	s.server.register <- client
	go client.writePump()
	go client.readPump(s.server)
}

func (s *service) SendAll(msg message) {
	s.server.mu.Lock()
	defer s.server.mu.Unlock()
	for _, client := range s.server.clients {
		client.send <- msg
	}
}

func (s *service) SendWithClientId(clientId string, msg message) error {
	s.server.mu.Lock()
	defer s.server.mu.Unlock()
	client := s.server.clients[clientId]

	if client == nil {
		slog.Error("No client found for given id", "clientId", clientId, "message", msg, "error", NoClientFoundErr)
		return NoClientFoundErr
	}

	client.send <- msg
	return nil
}
