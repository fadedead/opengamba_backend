package rewards

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /user/{id}", h.GetRewardByUserID)
	return mux
}

func (h *Handler) GetRewardByUserID(res http.ResponseWriter, req *http.Request) {
	userIdStr := req.PathValue("id")
	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		http.Error(res, "Invalid user ID", http.StatusBadRequest)
		return
	}

	reward, err := h.service.GetByUserId(uint(userId))
	if err != nil {
		http.Error(res, "Reward not found for user", http.StatusNotFound)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(reward)
}
