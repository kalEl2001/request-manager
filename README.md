# Request Manager

You need to set environment variable for:
- `LOGSTASH_URL`: TCP connection to logstash (e.g. `localhost:5000`)
- `RABBITMQ_URL`: AQMP connection to RabbitMQ (e.g. `aqmp://guest:guest@localhost:5672`)
- `OSMIUM_DOWNLOAD_FOLDER`: folder that will be used to store all request result

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


Testing TA
