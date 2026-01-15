package websocket

import (
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", h.SetupWebsocketConnectionForUser)
	return mux
}

func (h *Handler) SetupWebsocketConnectionForUser(res http.ResponseWriter, req *http.Request) {
	userId := req.URL.Query().Get("userId")
	conn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		return
	}

	h.service.registerClientConnection(userId, conn)
}
