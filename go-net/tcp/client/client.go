package main

import (
	"encoding/json"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Message struct {
	Msg string
}

var address = flag.String("address", ":8080", "")

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *address)
	if err != nil {
		log.Fatalf("Dial Server %s error: %v \n", *address, err)
	}
	defer conn.Close()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		conn.Close()
	}()
	encoder := json.NewEncoder(conn)
	for {
		msg := &Message{
			Msg: "Client Request",
		}
		if err := encoder.Encode(msg); err != nil {
			log.Printf("Send Message Error: %v \n", err)
			return
		}
		time.Sleep(time.Second * 5)
	}
}
