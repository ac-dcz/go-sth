package main

import (
	"crypto/tls"
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
	flag.Parse()
	serverCrt, err := tls.LoadX509KeyPair("../x509/server.crt", "../x509/server.key")
	if err != nil {
		log.Fatalf("Load Certificate error: %v \n", err)
	}
	tlsConf := tls.Config{
		Certificates: []tls.Certificate{serverCrt},
	}

	lis, err := net.Listen("tcp", *address)
	if err != nil {
		log.Fatalf("Listen %s error: %v \n", *address, err)
	}
	defer lis.Close()

	tlsLis := tls.NewListener(lis, &tlsConf) // 在TCP协议之上，加上一层TLS协议

	for {
		conn, err := tlsLis.Accept()
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
