import { Kafka } from 'kafkajs'
import dotenv from 'dotenv';
import fs from 'fs';
import readline from 'readline';
import path from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

dotenv.config();

const kafka = new Kafka({
  clientId: 'node-producer',
  brokers: ['kafka:9092'],
})

const producer = kafka.producer()

await producer.connect()

const sendOrders = async () => {
  const filePath = path.resolve(__dirname, '../../orders.ndjson');
  const fileStream = fs.createReadStream(filePath);

  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity,
  });

  for await (const line of rl) {
    if (!line.trim()) continue;

    try {
      const order = JSON.parse(line);
      await producer.send({
        topic: 'orders',
        messages: [
          { key: order.orderId, value: JSON.stringify(order) },
        ],
      });
      console.log(`✅ Sent: ${order.eventType} for ${order.orderId}`);
    } catch (err) {
      console.error(`❌ Failed to send message:`, err);
    }
  }

  await producer.disconnect();
};

sendOrders().catch(console.error);