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

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "kafka:9092"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

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

		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(key),
			Value:          orderJSON,
		}, nil)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		panic(err)
	}

	for p.Flush(10000) > 0 {
		fmt.Print("Still waiting to flush outstanding messages\n")
	}
}