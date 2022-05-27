package main

func main() {
    initDBConnection()
    migrateDB()
    
    initLogger()
    initRabbitMQ()

    readMessage()

    closeRabbitMQ()
}