package main

import (
	"encoding/json"
	"os"
    "strconv"
	amqp "github.com/rabbitmq/amqp091-go"
)

var rabbitmq *amqp.Connection
var rabbitChannel *amqp.Channel

func initRabbitMQ() {
	RABBITMQ_URL := os.Getenv("RABBITMQ_URL")
	if len(RABBITMQ_URL) == 0 {
		RABBITMQ_URL = "amqp://localhost:5672"
	}
	conn, err := amqp.Dial(RABBITMQ_URL)
	failLog(err, "Failed to connect to RabbitMQ")
	infoLog("Successfully connect to RabbitMQ", nil)

	rabbitmq = conn

    ch, err := rabbitmq.Channel()
    failLog(err, "Failed to open a channel")
    infoLog("Successfully open a channel in RabbitMQ", nil)

    rabbitChannel = ch
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
    ch := rabbitChannel

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

func publishMessage(route string, body map[string]interface{}, corrId uint) {
    ch := rabbitChannel
    bodyJson, err := json.Marshal(body)
    errorLog(err, "Failed to convert map to json")

    err = ch.Publish(
        "",     // exchange
        route, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing {
            ContentType: "application/json",
            replyTo: "servicemanager_queue",
            CorrelationId: strconv.FormatUint(uint64(corrId), 10),
            Body:        bodyJson,
        },
    )

    errorLog(err, "Failed to publish message")
}

func getDownloadFolder() string {
    ret := os.Getenv("OSMIUM_DOWNLOAD_FOLDER")
    if len(ret) == 0 {
        return "/osmium/result/"
    }
    return ret
}

func createDownloadJobMessage(slug string, link string, fileId uint) {
    body := map[string]interface{}{
        "url": link,
        "folder_out": getDownloadFolder() + slug,
    }
    publishMessage("downloader_queue", body, fileId)
}

func createCompressJobMessage(slug string, reqId uint) {
    body := map[string]interface{}{
        "folder": getDownloadFolder() + slug,
    }
    publishMessage("compressor_queue", body, reqId)
}

func updateStatusFileProvider(id uint, queryType string, data string) {
    body := map[string]interface{}{
        "query_type": queryType,
        "id": id,
    }
    if queryType == "update_progress" {
        body["progress"] = data
    } else if queryType == "update_url" {
        body["url"] = data
    } else if queryType == "create" {
        body["user_code"] = data
    }
    publishMessage("provider_queue", body, 0)
}
