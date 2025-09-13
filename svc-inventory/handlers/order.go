package handlers

import (
	"encoding/json"
	"log"

	"svc-inventory/dto"

	"github.com/segmentio/kafka-go"
)

func HandleOrderCreated(message kafka.Message) error {
	log.Printf("Processing message: %s", string(message.Value))
	var event dto.OrderCreatedEvent
	if err := json.Unmarshal(message.Value, &event); err != nil {
		return err
	}

	log.Printf("Processing order: %+v", event)
	log.Printf("📦 Checking stock for item: %s (qty: %d)", event.Item, event.Quantity)
	return nil
}
