module github.com/pramineni01/go-proto-exercise/server

go 1.13

replace github.com/pramineni01/go-proto-exercise/proto => ../../proto

require (
	github.com/golang/protobuf v1.3.5
	github.com/pramineni01/go-proto-exercise/proto v1.0.0
	golang.org/x/net v0.0.0-20200319234117-63522dbf7eec // indirect
	google.golang.org/genproto v0.0.0-20200319113533-08878b785e9c // indirect
	google.golang.org/grpc v1.28.0
)
