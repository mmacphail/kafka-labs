import json

with open('../../orders.ndjson', 'r') as file:
    for line in file:
        if line.strip():
            order = json.loads(line)
            order_id = order.get('orderId')
            print(f"Sending order {order_id}")