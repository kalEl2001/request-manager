package main

import (
	"encoding/json"
	"os"
    "strconv"
	amqp "github.com/rabbitmq/amqp091-go"
)

var rabbitmq *amqp.Connection

func initRabbitMQ() {
	RABBITMQ_URL := os.Getenv("RABBITMQ_URL")
	if len(RABBITMQ_URL) == 0 {
		RABBITMQ_URL = "amqp://osmium:osmium12345678@rabbitmq.faishol.net:5672"
	}
	conn, err := amqp.Dial(RABBITMQ_URL)
	failLog(err, "Failed to connect to RabbitMQ")
	infoLog("Successfully connect to RabbitMQ", nil)

	rabbitmq = conn
}

func closeRabbitMQ() {
	rabbitmq.Close()
    infoLog("Successfully close RabbitMQ connection", nil)
}

func parseMessageBody(msg amqp.Delivery) (map[string]interface{}) {
	var bodyJson map[string]interface{}
    err := json.Unmarshal(msg.Body, &bodyJson)
    errorLog(err, "Failed to parse message body")
    return bodyJson
}

func getRequestType(msg amqp.Delivery) string {
	bodyJson := parseMessageBody(msg)
	if val, ok := bodyJson["query_type"]; ok {
		return val.(string)
	}
	return ""
}

func readMessage() {
    ch, err := rabbitmq.Channel()
    failLog(err, "Failed to open a channel")
    infoLog("Successfully open a channel in RabbitMQ", nil)

    queue, err := ch.QueueDeclare(
        "servicemanager_queue",
        false,
        false,
        false,
        false,
        nil,
    )
    failLog(err, "Failed to declare a queue")
    infoLog("Successfully declare queue in RabbitMQ", nil)

    msgs, err := ch.Consume(
        queue.Name, // queue
        "",     // consumer
        true,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    failLog(err, "Failed to register a consumer")
    infoLog("Ready to receive messages", nil)

    for msg := range msgs {
        requestType := getRequestType(msg)
        if requestType == "create" {
            createRequest(parseMessageBody(msg))
        } else if requestType == "download" {
            corrId, err := strconv.Atoi(msg.CorrelationId)
            errorLog(err, "Cannot convert correlationId to integer")
            downloadResponse(corrId)
        } else if requestType == "compress" {
            corrId, err := strconv.Atoi(msg.CorrelationId)
            errorLog(err, "Cannot convert correlationId to integer")
            compressResponse(corrId, parseMessageBody(msg))
        } else {
			warningLog("Receive unknown request type", nil)
		}
    }
}