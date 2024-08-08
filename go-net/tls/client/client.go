package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"log"
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
	crt, err := os.ReadFile("../x509/ca.crt")
	if err != nil {
		log.Fatalf("Read Root Certificate Error: %v \n", err)
	}
	credsPool := x509.NewCertPool()
	credsPool.AppendCertsFromPEM(crt)
	tlsConf := tls.Config{
		ServerName: "www.home.dcz.com",
		RootCAs:    credsPool,
	}

	conn, err := tls.Dial("tcp", *address, &tlsConf)
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
