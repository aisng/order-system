package handlers

import (
	"encoding/json"
	"fmt"
	"log"

	"svc-inventory/dto"
	"svc-inventory/persistence"

	"github.com/segmentio/kafka-go"
)

func HandleOrderCreated(message kafka.Message) error {
	log.Printf("Processing message: %s", string(message.Value))

	var event dto.OrderCreatedEvent
	if err := json.Unmarshal(message.Value, &event); err != nil {
		return err
	}

	repo := persistence.NewRepository()
	status := "pending"
	orderID, err := repo.CreateOrder(event.ItemID, event.BuyerAddress, status)
	if err != nil {
		return fmt.Errorf("Failed to create order: %s", err.Error())
	}

	log.Printf("created order id: %d", orderID)
	return nil
}
