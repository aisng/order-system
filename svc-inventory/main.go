package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
	"svc-inventory/async"
	"svc-inventory/handlers"
	"syscall"
)

func main() {
	kafkaBroker := os.Getenv("KAFKA_BROKER")
	if kafkaBroker == "" {
		kafkaBroker = "localhost:9092"
	}

	log.Printf("Starting svc-inventory...")
	log.Printf("Connecting to Kafka at: %s", kafkaBroker)

	brokers := strings.Split(kafkaBroker, ",")
	consumer := async.NewConsumer(brokers, "order.created", "inventory-group")
	defer consumer.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		log.Println("Shutdown signal received")
		cancel()
	}()

	consumer.Start(ctx, handlers.HandleOrderCreated)
}
