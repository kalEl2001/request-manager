FROM golang:1.18-alpine AS build

WORKDIR /app
RUN apk update \
    && apk --no-cache --update add build-base bind-tools
    
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build -o request-manager -v

ENTRYPOINT ["./request-manager"]