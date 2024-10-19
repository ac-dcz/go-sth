package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	Topic    = "example-gosth-test"
	Nums     = 20
	Interval = time.Second * 3
)

func main() {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{"127.0.0.1:9092"},
		Topic:        Topic,
		RequiredAcks: int(kafka.RequireAll),
		Logger:       log.New(os.Stdout, "[kafka]", log.LstdFlags),
		ErrorLogger:  log.New(os.Stderr, "[kafka]", log.LstdFlags),
	})

	w.AllowAutoTopicCreation = true
	defer w.Close()

	for i := 0; i < Nums; i++ {
		msg := kafka.Message{
			Key:   []byte(Topic),
			Value: []byte(fmt.Sprintf("test message %d %d", time.Now().Unix(), i)),
		}

		if err := w.WriteMessages(context.Background(), msg); err != nil {
			if !errors.Is(err, kafka.LeaderNotAvailable) && !errors.Is(err, context.DeadlineExceeded) {
				log.Println(err)
				break
			}
		}
		time.Sleep(Interval)
	}
}
