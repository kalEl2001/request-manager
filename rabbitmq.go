package main

import (
	"os"
	amqp "github.com/rabbitmq/amqp091-go"
)

func initRabbitMQ() *amqp.Connection {
	RABBITMQ_URL := os.Getenv("RABBITMQ_URL")
	if len(RABBITMQ_URL) == 0 {
		RABBITMQ_URL = "amqp://osmium:osmium12345678@rabbitmq.faishol.net:5672"
	}
	conn, err := amqp.Dial(RABBITMQ_URL)
	if err != nil {
		field := map[string]interface{}{
			"rabbitmq_url": RABBITMQ_URL,
			"error": err,
		}
		createLog(nil, "Panic", "Failed to connect to RabbitMQ", field)
	}
	createLog(nil, "Info", "Successfully to connect to RabbitMQ", nil)

	return conn
}

