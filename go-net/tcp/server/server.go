package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net"
)

type Message struct {
	Msg string
}

var address = flag.String("address", ":8080", "")

func main() {
	lis, err := net.Listen("tcp", *address)
	if err != nil {
		log.Fatalf("Listen %s error: %v \n", *address, err)
	}
	defer lis.Close()
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("Accept Error: %v \n", err)
		}
		go func() {
			defer conn.Close()
			decoder := json.NewDecoder(conn)
			for {
				msg := &Message{}
				if err := decoder.Decode(msg); err != nil {
					if err != io.EOF {
						log.Printf("Decoder Message error: %v \n", err)
						return
					}
				}
				log.Printf("Receive Client Message: %s \n", msg.Msg)
			}

		}()
	}

}
