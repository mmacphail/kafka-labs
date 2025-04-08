from kafka import KafkaProducer
import json

producer = KafkaProducer(
    bootstrap_servers='kafka:9092',
    key_serializer=lambda k: json.dumps(k).encode('utf-8'),
    value_serializer=lambda v: json.dumps(v).encode('utf-8')
)

futures = []

with open('../../orders.ndjson', 'r') as file:
    for line in file:
        if line.strip():
            order = json.loads(line)
            order_id = order.get('orderId')
            future = producer.send('orders', key=order_id, value=order)
            futures.append(future)

for future in futures:
    try:
        metadata = future.get(timeout=10)
        print(f"Message sent to partition {metadata.partition} at offset {metadata.offset}")
    except Exception as e:
        print(f"Failed to send message: {e}")

producer.flush()
producer.close()