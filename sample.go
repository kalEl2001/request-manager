package main
// package main
// import (
//         "github.com/bshuster-repo/logrus-logstash-hook"
//         "github.com/sirupsen/logrus"
//         "net"
// )
// func main() {
//         log := logrus.New()
//         conn, err := net.Dial("tcp", "34.132.46.126:5000") 
//         if err != nil {
//                 log.Fatal(err)
//         }
//         hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{"type": "logs"}))
//         log.Hooks.Add(hook)
//         ctx := log.WithFields(logrus.Fields{
//                 "haha": "logs",
//         })
//         ctx.Info("Info Message")
//         ctx.Error("Error Message")
//         ctx.Warn("Warn Message")
//         ctx.Fatal("Fatal Message")
// }

// func failOnError(err error, msg string) {
//     if err != nil {
//         log.Fatalf("%s: %s", msg, err)
//     }
// }

// func main() {
//     conn, err := amqp.Dial("amqp://osmium:osmium12345678@osmium.faishol.net:5672/")
//     failOnError(err, "Failed to connect to RabbitMQ")
//     defer conn.Close()

//     ch, err := conn.Channel()
//     failOnError(err, "Failed to open a channel")
//     defer ch.Close()

//     q, err := ch.QueueDeclare(
//         "hello", // name
//         false,   // durable
//         false,   // delete when unused
//         false,   // exclusive
//         false,   // no-wait
//         nil,     // arguments
//     )
//     failOnError(err, "Failed to declare a queue")
    
//     body := "Hello World!"
//     err = ch.Publish(
//         "",     // exchange
//         q.Name, // routing key
//         false,  // mandatory
//         false,  // immediate
//         amqp.Publishing {
//           ContentType: "text/plain",
//           Body:        []byte(body),
//         })
//     failOnError(err, "Failed to publish a message")
// }

// func openRabbitConnection() (*Connection, error) {
//     AQMP_URL := "amqp://osmim:osmium12345678@osmium.faishol.net:5672/"
//     conn, err := amqp.Dial(AQMP_URL)
//     if err != nil {
//         nil, fmt.Errorf("[%s] - %s", err, AQMP_URL)
//     }
//     defer conn.Close()


// }