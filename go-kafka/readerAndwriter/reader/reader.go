package main

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	Address  = "127.0.0.1:9092"
	Topic    = "example-gosth-test"
	Nums     = 20
	Interval = time.Second * 2
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{Address},
		Topic:       Topic,
		Logger:      log.New(os.Stdout, "[kafka]", log.LstdFlags),
		ErrorLogger: log.New(os.Stderr, "[kafka]", log.LstdFlags),
	})
	r.SetOffset(kafka.LastOffset) //从最新的位置开始消费
	defer r.Close()
	for i := 0; i < Nums; i++ {
		msg, err := r.FetchMessage(context.Background())
		if err != nil {
			if !errors.Is(err, kafka.LeaderNotAvailable) && !errors.Is(err, context.DeadlineExceeded) {
				log.Println(err)
				break
			}
		}
		log.Println("============================================")
		log.Printf("Value: %s P: %d Off: %d \n", msg.Value, msg.Partition, msg.Offset)
		for _, h := range msg.Headers {
			log.Printf("Key: %s Value: %s \n", h.Key, h.Value)
		}
		time.Sleep(Interval)
	}

}
