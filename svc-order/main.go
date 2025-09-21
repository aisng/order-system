package main

import (
	"log"
	"net/http"
	"svc-order/async"
	"svc-order/handlers"

	"github.com/gorilla/mux"
)

func main() {
	log.Printf("Starting svc-order...")

	producer := async.NewProducer(async.GetBrokers())

	orderHander := handlers.NewOrderHandler(producer)

	router := mux.NewRouter()

	router.HandleFunc("/orders", orderHander.CreateOrder).Methods("POST")

	log.Printf("svc-order ready on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
