package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"svc-order/async"
	"svc-order/handlers"

	"github.com/gorilla/mux"
)

func main() {
	kafkaBroker := os.Getenv("KAFKA_BROKER")
	if kafkaBroker == "" {
		kafkaBroker = "localhost:9092"
	}

	log.Printf("Starting svc-order...")
	log.Printf("Connecting to Kafka at: %s", kafkaBroker)

	brokers := strings.Split(kafkaBroker, ",")
	producer := async.NewProducer(brokers, "order.created")
	defer producer.Close()

	orderHander := handlers.NewOrderHandler(producer)

	router := mux.NewRouter()

	router.HandleFunc("/orders", orderHander.CreateOrder).Methods("POST")

	log.Printf("svc-order ready on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
