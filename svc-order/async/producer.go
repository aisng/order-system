package async

import (
	"context"
	"encoding/json"
	"fmt"
	"svc-order/dto"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	Writer *kafka.Writer
}

func NewProducer(brokers []string, topic string) *Producer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
	})

	return &Producer{
		Writer: writer,
	}
}

func (p *Producer) Close() error {
	return p.Writer.Close()
}

func (p *Producer) PublishOrderCreated(order dto.Order) error {
	event := dto.OrderCreatedEvent{
		OrderID:      order.ID,
		BuyerAddress: order.BuyerAddress,
		Item:         order.Item,
		Quantity:     order.Quantity,
		Amount:       order.Amount,
	}

	eventJSON, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("Error marshalling event: %v", err)
	}

	err = p.Writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(order.ID),
			Value: eventJSON,
		},
	)

	if err != nil {
		return fmt.Errorf("Error sending Kafka message: %v", err)
	}

	return nil
}
