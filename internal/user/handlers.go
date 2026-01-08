package user

import (
	"net/http"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

func (handler *handler) Routes() http.Handler {
	mux := http.NewServeMux()

	return mux
}

