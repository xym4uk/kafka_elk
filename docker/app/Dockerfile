FROM golang:1.20-buster

WORKDIR /app

RUN go version

ENV GOPATH=/

COPY ./ ./

RUN go clean -modcache
RUN go mod tidy

RUN go build -o consumer /app/cmd/consumer/consumer.go
RUN go build -o producer /app/cmd/producer/producer.go