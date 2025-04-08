# Hints

## Adding the KafkaJS library

```sh
npm install kafkajs
```

### Creating a producer

```js
const kafka = new Kafka({
  clientId: 'node-producer',
  brokers: ['kafka:9092'],
})

const producer = kafka.producer()

await producer.connect()
```

### Send the messages

```js
await producer.send({
  topic: 'orders',
  messages: [
    { key: order.orderId, value: JSON.stringify(order) },
  ],
});
```

### Disconnect the producer

```js
await producer.disconnect();
```

### Run the producer

```js
node index.js
```