package http

import (
	"encoding/json"
	"hexagonal_in_go_way/internal/application"
	"hexagonal_in_go_way/internal/domain/entities"
	"net/http"
)

type OrderHandler struct {
	service *application.OrderService
}

func NewOrderHandler(service *application.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order entities.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateOrder(r.Context(), &order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
