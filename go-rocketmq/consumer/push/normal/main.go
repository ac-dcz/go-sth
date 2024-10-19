package main

import (
	"context"
	"log"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

const (
	Address   = "127.0.0.1:9876"
	GroupName = "CGN_example_test-3"
	Topic     = "example_gosth_test"
	Nums      = 20
	Interval  = time.Second * 2
)

func main() {
	r, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer(primitive.NamesrvAddr{Address}),
		consumer.WithGroupName(GroupName),
		consumer.WithRetry(2),
		consumer.WithConsumeMessageBatchMaxSize(1),
		consumer.WithConsumeGoroutineNums(1),
	)
	if err != nil {
		log.Fatalln(err)
	}
	selector := consumer.MessageSelector{
		Type:       consumer.TAG,
		Expression: "normal",
	}
	if err := r.Subscribe(Topic, selector, consumeMessage); err != nil {
		log.Fatalln(err)
	}
	if err := r.Start(); err != nil {
		log.Fatalln(err)
	}
	defer r.Shutdown()
	time.Sleep(time.Second * 30)
}

func consumeMessage(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	log.Println("=============================================")
	log.Printf("Message Len = %d \n", len(msgs))
	for _, msg := range msgs {
		log.Printf("MsgId = %s, OffMsgId = %s, QueueOff = %d \n", msg.MsgId, msg.OffsetMsgId, msg.QueueOffset)
		log.Printf("Queue = %s \n", msg.Queue.String())
		log.Printf("Msg Body = %s Tag = %s\n", msg.Body, msg.GetTags())
	}
	time.Sleep(Interval)
	return consumer.ConsumeSuccess, nil
}
