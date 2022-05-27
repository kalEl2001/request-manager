# Request Manager

You need to set environment variable for:
- `LOGSTASH_URL`: TCP connection to logstash (e.g. `localhost:5000`)
- `RABBITMQ_URL`: AQMP connection to RabbitMQ (e.g. `aqmp://guest:guest@localhost:5672`)

## Install module
```
go mod download
```

## Build
```
go build .
```

## Run
```
go run .
```
