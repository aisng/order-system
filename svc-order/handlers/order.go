package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"svc-order/async"
	"svc-order/dto"
	"time"
)

type OrderHandler struct {
	Producer *async.Producer
}

func NewOrderHandler(producer *async.Producer) *OrderHandler {
	return &OrderHandler{
		Producer: producer,
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order dto.Order

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		log.Printf("Error decoding orderL %v", err)
		http.Error(w, "Invalid request JSON", http.StatusBadRequest)
		return
	}

	order.ID = fmt.Sprintf("order-%d", time.Now().Unix())
	order.Status = "pending"

	log.Printf("Creating order: %+v", order)

	if err := h.Producer.PublishOrderCreated(order); err != nil {
		http.Error(w, fmt.Sprintf("Failed to process order: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
