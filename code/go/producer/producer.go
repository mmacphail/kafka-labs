package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Order struct {
	OrderID     string  `json:"orderId"`
	EventType   string  `json:"eventType"`
	Timestamp   string  `json:"timestamp"`
	Service     string  `json:"service"`
	Customer    string  `json:"customer,omitempty"`
	Location    string  `json:"location,omitempty"`
}

func main() {
	file, err := os.Open("../../orders.ndjson")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	topic := "orders"

	for scanner.Scan() {
		line := scanner.Text()
		var order Order
		
		if err := json.Unmarshal([]byte(line), &order); err != nil {
			fmt.Printf("Error parsing order: %v\n", err)
			panic(err)
		}
		key := order.OrderID

		orderJSON, err := json.Marshal(order)
		if err != nil {
			fmt.Printf("Error marshaling order: %v\n", err)
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		panic(err)
	}
}