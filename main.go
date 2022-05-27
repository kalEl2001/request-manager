package main

func main() {
    migrateDB()
    
    initLogger()
    initRabbitMQ()

    readMessage()

    closeRabbitMQ()
}