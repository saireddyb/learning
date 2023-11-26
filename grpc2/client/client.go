package main

import (
	"context"
	"fmt"
	"log"
	"workshop/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc, err := grpc.Dial("localhost:8080", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := pb.NewGreetingServiceClient(cc)
	request := &pb.GreetingServiceRequest{Name: "Gophers"}

	resp, err := client.Greeting(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive response => %s ", resp.Message)
}