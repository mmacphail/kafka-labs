# Hints

## Adding the Confluent Kafka Go Library

```go
require (
    github.com/confluentinc/confluent-kafka-go/v2 v2.8.0
)
```

### Creating a producer with properties

```go
p, err := kafka.NewProducer(&kafka.ConfigMap{
    "bootstrap.servers": "kafka:9092",
    "acks":              "all",
})

defer p.Close()
```

### Handle producer events

```go
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
```

### Send the messages

```go
p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(key),
			Value:          orderJSON,
		}, nil)
```

### Flush the producer

```go
for p.Flush(10000) > 0 {
  fmt.Print("Still waiting to flush outstanding messages\n")
}
```

### Run the producer

```go
go run producer.go
```