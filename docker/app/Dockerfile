FROM golang:1.19-buster

WORKDIR /app

RUN go version

ENV GOPATH=/

COPY ./go.mod ./

RUN go mod download

RUN go mod download github.com/confluentinc/confluent-kafka-go/v2

COPY ./cmd ./

COPY ./out ./out
