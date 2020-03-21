package main

import (
	"context"
	"reflect"
	"testing"
	"time"

	proto "github.com/pramineni01/go-proto-exercise/proto"
)

func Test_server_SayHello(t *testing.T) {
	type args struct {
		ctx context.Context
		req *proto.HelloRequest
	}
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Millisecond)
	tests := []struct {
		name    string
		args    args
		wantRes *proto.HelloResponse
		wantErr bool
	}{
		{
			name:    "Success",
			args:    args{ctx: ctx, req: &proto.HelloRequest{Name: "Name1"}},
			wantRes: &proto.HelloResponse{Greetings: "Hello Name1"},
			wantErr: false,
		},
	}

	srv := server{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := srv.SayHello(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.SayHello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("server.SayHello() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
