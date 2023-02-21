package main

import (
	"fmt"
	"github.com/bxcodec/faker/v4"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
	"os"
	"strconv"
)

type MyLog struct {
	Timestamp string `faker:"timestamp"`
	Log       string `faker:"sentence"`
}

func main() {
	args := os.Args
	if len(args) == 1 {
		log.Fatal("Не указано количество логов")
	}
	logsCount, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal("указано некорректное число")
	}

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:9093",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	topic := "test"
	myLog := MyLog{}

	for i := 0; i < logsCount; i++ {
		err = faker.FakeData(&myLog)
		if err != nil {
			log.Fatal("error occured", err)
		}

		logString := fmt.Sprintf("%s - %s", myLog.Timestamp, myLog.Log)
		err := producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(logString),
		}, nil)
		if err != nil {
			log.Fatal(err)
		}
	}

	producer.Flush(15 * 1000)
}
