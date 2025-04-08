# Hints

## Adding the kafka-python library

```sh
pip install kafka-python
```

### Creating a producer

```python
from kafka import KafkaProducer
producer = KafkaProducer(
    bootstrap_servers='kafka:9092',
    key_serializer=lambda k: json.dumps(k).encode('utf-8'),
    value_serializer=lambda v: json.dumps(v).encode('utf-8')
)
```

### Send the messages

```python
future = producer.send('orders', key=order_id, value=order)
```

## Make sure the messages have been sent

```python
for future in futures:
    try:
        metadata = future.get(timeout=10)
        print(f"Message sent to partition {metadata.partition} at offset {metadata.offset}")
    except Exception as e:
        print(f"Failed to send message: {e}")
```

### Disconnect the producer

```python
producer.flush() # Optional
producer.close()
```

### Run the producer

```js
python producer.py
```