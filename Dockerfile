FROM golang:1.18-alpine AS build

WORKDIR /app
RUN apk update \
    && apk --no-cache --update add build-base bind-tools
    
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o request-manager -v

COPY . .
ENTRYPOINT ["./request-manager"]