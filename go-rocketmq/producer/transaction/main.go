package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

// type TransactionListener interface {
// 	//  When send transactional prepare(half) message succeed, this method will be invoked to execute local transaction.
// 	ExecuteLocalTransaction(*Message) LocalTransactionState

// 	// When no response to prepare(half) message. broker will send check message to check the transaction status, and this
// 	// method will be invoked to get local transaction status.
// 	CheckLocalTransaction(*MessageExt) LocalTransactionState
// }

type SimpleTxListener struct {
	state sync.Map
}

func (sTx *SimpleTxListener) ExecuteLocalTransaction(msg *primitive.Message) primitive.LocalTransactionState {
	log.Println("ExecuteLocalTransaction")
	log.Printf("Topic: %s\t TransactionId: %s \n", msg.Topic, msg.TransactionId)
	state := rand.Intn(3)
	sTx.state.Store(msg.TransactionId, primitive.LocalTransactionState(state))

	return primitive.UnknowState //强制触发回查
}

func (sTx *SimpleTxListener) CheckLocalTransaction(msg *primitive.MessageExt) primitive.LocalTransactionState {
	log.Println("CheckLocalTransaction")
	log.Printf("Topic: %s\tQueueId: %d\tTransactionId: %s \n", msg.Topic, msg.Queue.QueueId, msg.TransactionId)
	data, ok := sTx.state.Load(msg.TransactionId)
	if !ok {
		log.Printf("unknown msg: %v, return Commit", msg)
		return primitive.CommitMessageState
	}
	state := data.(primitive.LocalTransactionState)
	switch state {
	case 0:
		log.Printf("checkLocalTransaction COMMIT_MESSAGE: %v\n", msg)
		return primitive.CommitMessageState
	case 1:
		log.Printf("checkLocalTransaction ROLLBACK_MESSAGE: %v\n", msg)
		return primitive.RollbackMessageState
	case 2:
		log.Printf("checkLocalTransaction unknown: %v\n", msg)
		return primitive.UnknowState
	default:
		log.Printf("checkLocalTransaction default COMMIT_MESSAGE: %v\n", msg)
		return primitive.CommitMessageState
	}
}

const (
	NameServer = "127.0.0.1:9876"
	Topic      = "TestTopic"
	GroupName  = "PGN-example-test-1"
	Nums       = 20
	Interval   = time.Second * 5
)

func main() {
	txw, err := rocketmq.NewTransactionProducer(
		&SimpleTxListener{},
		producer.WithNameServer(primitive.NamesrvAddr{NameServer}),
		producer.WithGroupName(GroupName),
		producer.WithRetry(2),
	)
	if err != nil {
		log.Fatalln(err)
	}
	if err := txw.Start(); err != nil {
		log.Fatalln(err)
	}
	defer txw.Shutdown()
	for i := 0; i < Nums; i++ {
		msg := primitive.Message{
			Topic: Topic,
			Body:  []byte(fmt.Sprintf("transaction message %d %d", time.Now().Unix(), i)),
		}
		result, err := txw.SendMessageInTransaction(context.Background(), &msg)
		if err != nil {
			log.Println(err)
		}
		log.Println("======================================")
		log.Printf("MsgID=%s,OffsetMsgID=%s,QueueOffset=%d\n", result.MsgID, result.OffsetMsgID, result.QueueOffset)
		log.Printf("Queue: %s", result.MessageQueue.String())

		time.Sleep(Interval)
	}
}
