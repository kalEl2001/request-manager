#
## Build
##
FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN apk update \
    && apk --no-cache --update add build-base \
    && go build -o /osmium-request-manager

##
## Deploy
##
FROM alpine

WORKDIR /app/osmium

COPY .env .
COPY --from=build /osmium-request-manager .
USER nonroot:nonroot

ENTRYPOINT ["./osmium-request-manager"]