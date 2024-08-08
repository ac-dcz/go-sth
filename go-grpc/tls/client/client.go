package main

import (
	"context"
	"flag"
	"go-grpc/tls/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var address = flag.String("address", ":8080", "")

func main() {
	flag.Parse()
	ca, err := credentials.NewClientTLSFromFile("../x509/ca.crt", "www.home.dcz.com")
	if err != nil {
		log.Fatalf("Create Credential Error: %v \n", err)
	}
	conn, err := grpc.NewClient(*address, grpc.WithTransportCredentials(ca))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	client := pb.NewHelloClient(conn)
	resp, err := client.SayHello(context.Background(), &pb.HelloReq{Name: "dcz"})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%s \n", resp.Msg)
}
