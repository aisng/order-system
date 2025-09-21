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
	Producer async.Producer
}

func NewOrderHandler(producer async.Producer) *OrderHandler {
	return &OrderHandler{
		Producer: producer,
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order dto.Order

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		log.Printf("Error decoding order: %v", err)
		http.Error(w, "Invalid request JSON", http.StatusBadRequest)
		return
	}

	// repo := persistence.NewRepository()
	// orderID, err := repo.CreateOrder(order.ItemID, order.BuyerAddress, order.Status)
	//
	// if err != nil {
	// 	log.Printf("failed to persist order: id %d: %v", orderID, err)
	// 	http.Error(w, fmt.Sprintf("failed to persist order: %v", err), http.StatusInternalServerError)
	// 	return
	// }

	orderID := fmt.Sprintf("order-%d", time.Now().Unix())
	orderEvent := dto.OrderEvent{
		Type:         "created",
		OrderID:      orderID,
		BuyerAddress: order.BuyerAddress,
		ItemID:       order.ItemID,
		Status:       "pending",
	}

	log.Printf("publishing event: %+v", orderEvent)

	if err := h.Producer.PublishEvent("order", orderID, orderEvent); err != nil {
		log.Printf("error publishing order event: %v", err)
		http.Error(w, fmt.Sprintf("Failed to process order: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
