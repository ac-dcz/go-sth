package main

import (
	"context"
	"flag"
	"fmt"
	"go-grpc/tls/pb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type HelloService struct {
	pb.UnimplementedHelloServer
}

func (h *HelloService) SayHello(ctx context.Context, r *pb.HelloReq) (*pb.HelloResp, error) {
	msg := fmt.Sprintf("Hello %s", r.Name)
	log.Printf("Call SayHello %s \n", r.Name)
	return &pb.HelloResp{
		Msg: msg,
	}, nil
}

var address = flag.String("address", ":8080", "")

func main() {
	flag.Parse()
	cred, err := credentials.NewServerTLSFromFile("../x509/server.crt", "../x509/server.key")
	if err != nil {
		log.Printf("Create Server Credential Error: %v \n", err)
	}
	s := grpc.NewServer(grpc.Creds(cred))
	pb.RegisterHelloServer(s, &HelloService{})

	lis, err := net.Listen("tcp", *address)
	if err != nil {
		log.Fatalf("Listen %s error: %v \n", *address, err)
	}

	s.Serve(lis)
}
