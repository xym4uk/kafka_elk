FROM golang:1.19-buster

RUN go version
ENV GOPATH=/

RUN go build -o app ./producer.go

CMD ["./app"]