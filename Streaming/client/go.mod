module github.com/pramineni01/go-proto-exercise/client

go 1.13

replace github.com/pramineni01/go-proto-exercise/proto => ../../proto

require (
	github.com/pramineni01/go-proto-exercise/proto v1.0.0
	google.golang.org/grpc v1.28.0 // indirect
)
