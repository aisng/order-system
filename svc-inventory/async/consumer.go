package async

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	Reader *kafka.Reader
}

func NewConsumer(brokers []string, topic, groupID string) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupID,
		MaxBytes: 10e6,
	})

	return &Consumer{
		Reader: reader,
	}
}

func (c *Consumer) ProcessMessages(ctx context.Context, handler func(kafka.Message) error) {
	log.Printf("Starting consumer...")

	for {
		select {
		case <-ctx.Done():
			log.Println("Consumer stopped")
			return
		default:
			msgCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
			message, err := c.Reader.ReadMessage(msgCtx)
			cancel()
			if err != nil {
				if errors.Is(err, context.DeadlineExceeded) || strings.Contains(err.Error(), "deadline exceeded") {
					continue
				}
				log.Printf("Error reading message: %v", err)
				continue
			}

			log.Printf("Received message: key=%s, offset=%d",
				string(message.Key), message.Offset)
			if err := handler(message); err != nil {
				log.Printf("Error proccessing message: %v", err)
			}
		}
	}
}
func (c *Consumer) Close() error {
	log.Println("Closing consumer...")
	return c.Reader.Close()
}
