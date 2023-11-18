package main

import (
	"context"
	"fmt"
	hello "learning/grpc-hello-world/proto"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	hello.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	return &hello.HelloResponse{Message: "hello, " + req.Name + "!"}, nil
}

func main(){
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	hello.RegisterGreeterServer(s, &server{})
	fmt.Println("Port listening on : 50051")
	s.Serve(lis)
}