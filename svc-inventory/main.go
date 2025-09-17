package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
	"svc-inventory/async"
	"svc-inventory/handlers"
	"sync"
	"syscall"
)

func main() {
	log.Printf("Starting svc-inventory...")

	ctx, cancel := setupShudown()
	defer cancel()

	consumerWG := startConsumer(ctx)
	waitForShudownSignal()
	log.Println("shudown signal received, stopping consumer")

	cancel()

	consumerWG.Wait()
}

func setupShudown() (context.Context, context.CancelFunc) {
	return context.WithCancel(context.Background())
}

func waitForShudownSignal() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}

func startConsumer(ctx context.Context) *sync.WaitGroup {
	brokers := strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
	log.Printf("brokers from env: %s", brokers)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		consumer := async.NewConsumer(
			brokers,
			"order.created",
			"inventory-group",
		)
		defer consumer.Close()

		log.Println("Starting Kafka consumer...")
		consumer.Start(ctx, handlers.HandleOrderCreated)
		log.Println("Consumer stopped")
	}()

	return &wg
}
