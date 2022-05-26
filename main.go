package main

import (
    amqp "github.com/rabbitmq/amqp091-go"
    "github.com/sirupsen/logrus"
)

var logger *logrus.Logger
var rabbitmq *amqp.Connection

func failOnError(err error, msg string) {
    if err != nil {
        field := map[string]interface{}{
            "error_msg": err,
        }
        createLog(logger, "Info", msg, field)
    }
}

func main() {
    MigrateDB()
    logger = initLogger()
    rabbitmq = initRabbitMQ()

    readMessage()

    defer rabbitmq.Close()
    createLog(logger, "Info", "Successfully close RabbitMQ connection", nil)
}

func sendDownloadRequest(req_id int, link string) {

}

func sendCompressRequest(req_id int) {

}

func parseRequest() {

}

func readMessage() {
    ch, err := rabbitmq.Channel()
    failOnError(err, "Failed to open a channel")

    queue, err := ch.QueueDeclare(
        "servicemanager_queue",
        false,
        false,
        true,
        false,
        nil,
    )
    failOnError(err, "Failed to declare a queue")

    msgs, err := ch.Consume(
        queue.Name, // queue
        "",     // consumer
        true,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    failOnError(err, "Failed to register a consumer")

    for msg := range msgs {
        
    }
}