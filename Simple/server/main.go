package main

import (
	"context"
	"log"
	"net"

	proto "github.com/pramineni01/go-proto-exercise/proto"
	"google.golang.org/grpc"
)

type server struct {
	srv *proto.UnimplementedGreeterServer
}

func (srv *server) SayHello(ctx context.Context, req *proto.HelloRequest) (res *proto.HelloResponse, err error) {
	log.Print("Request received. Name:", req.GetName())
	return &proto.HelloResponse{Greetings: "Hello " + req.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Error opening port to listen. Err: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error starting the server. Err: %v", err)
	}
}
