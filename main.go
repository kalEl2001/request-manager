package main

import "github.com/joho/godotenv"

func main() {
    godotenv.Load()
    initLogger()

    initDBConnection()
    migrateDB()
    
    initRabbitMQ()
    readMessage()
    closeRabbitMQ()
}