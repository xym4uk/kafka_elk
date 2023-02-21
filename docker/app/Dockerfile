FROM golang:1.20-buster

WORKDIR /app

RUN go version

ENV GOPATH=/

COPY ./go.mod ./

RUN go mod download

COPY ./cmd ./cmd

RUN go build -o consumer /app/cmd/consumer/consumer.go
RUN go build -o producer /app/cmd/producer/producer.go