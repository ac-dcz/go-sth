package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

const (
	Address   = "127.0.0.1:9876"
	GroupName = "PGN_example_test-1"
	Topic     = "example_gosth_test"
	Nums      = 20
	Interval  = time.Second * 3
)

func main() {
	w, err := rocketmq.NewProducer(
		producer.WithNameServer(primitive.NamesrvAddr{Address}),
		producer.WithGroupName(GroupName),
		producer.WithRetry(2),
		producer.WithQueueSelector(producer.NewHashQueueSelector()), //设置根据Key hash 来选择Queue
	)
	if err != nil {
		log.Fatalln(err)
	}
	if err := w.Start(); err != nil {
		log.Fatalln(err)
	}
	defer w.Shutdown()
	ctx := context.TODO()
	for i := 0; i < Nums; i++ {
		msg := primitive.Message{
			Topic: Topic,
			Body:  []byte(fmt.Sprintf("normal message %d %d", time.Now().Unix(), i)),
		}
		msg.WithTag("order")
		msg.WithShardingKey("order") //同一个key的消息，存放在同一个Queue，实现消息的顺序存放
		//如果是发送多个消息，这些消息应该属于同一个Topic
		result, err := w.SendSync(ctx, &msg)
		if err != nil {
			log.Println(err)
		}
		log.Println("======================================")
		log.Printf("MsgID=%s,OffsetMsgID=%s,QueueOffset=%d\n", result.MsgID, result.OffsetMsgID, result.QueueOffset)
		log.Printf("Queue: %s", result.MessageQueue.String())

		time.Sleep(Interval)
	}
}
