package main

func main() {
    initLogger()

    initDBConnection()
    migrateDB()
    
    initRabbitMQ()

    readMessage()

    closeRabbitMQ()
}