package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"

	"github.com/pramineni01/go-proto-exercise/proto"
	"google.golang.org/grpc"
)

type server struct {
	srv *proto.UnimplementedMultiGreeterServer
}

func (s *server) GreetEach(stream proto.MultiGreeter_GreetEachServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		g := fmt.Sprintf("Hello %s", strings.ToUpper(in.Name))
		if err := stream.Send(&proto.Response{Greeting: g}); err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
}

func main() {
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Error opening port to listen. Err: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterMultiGreeterServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error starting the server. Err: %v", err)
	}
}
