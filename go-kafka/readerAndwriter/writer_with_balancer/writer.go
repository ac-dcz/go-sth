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

// type Balancer interface {
// 	// Balance receives a message and a set of available partitions and
// 	// returns the partition number that the message should be routed to.
// 	//
// 	// An application should refrain from using a balancer to manage multiple
// 	// sets of partitions (from different topics for examples), use one balancer
// 	// instance for each partition set, so the balancer can detect when the
// 	// partitions change and assume that the kafka topic has been rebalanced.
// 	Balance(msg Message, partitions ...int) (partition int)
// }

type MyRoundRobin struct {
	nextIndex int
}

func (mrb *MyRoundRobin) Balance(msg kafka.Message, partitions ...int) (partition int) {
	log.Printf("MyRoundRoBin")
	ind := mrb.nextIndex % len(partitions)
	mrb.nextIndex++
	return partitions[ind]
}

func main() {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{"127.0.0.1:9092"},
		Topic:        Topic,
		Balancer:     &MyRoundRobin{},
		RequiredAcks: int(kafka.RequireAll),
		Async:        false,
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
