package main

import (
	"context"
	"log"
	"time"

	"github.com/pramineni01/go-proto-exercise/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Error while connecting to server. Err: %v", err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("Error while communicating wiht server. Error: %v", err)
	}
	log.Println("Server response: ", r.GetGreetings())
}
