package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/pramineni01/go-proto-exercise/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Error while connecting to server. Err: %v", err)
	}
	defer conn.Close()
	c := proto.NewMultiGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	stream, err := c.GreetEach(ctx)
	if err != nil {
		log.Fatalf("Error while communicating wiht server. Error: %v", err)
	}

	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Printf("Server greets %s", in.Greeting)
		}
	}()

	names := []string{"Name1", "Name2", "Name3", "Name4", "Name5", "Name6", "Name7", "Name8", "Name9", "Name10"}

	for _, name := range names {
		if err := stream.Send(&proto.Request{Name: name}); err != nil {
			log.Fatalf("Failed to send name: %v", err)
		}
	}
	stream.CloseSend()
	<-waitc
}
