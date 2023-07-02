package main

import (
	"context"
	"hello_grpc/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	connType := insecure.NewCredentials()
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(connType))
	if err != nil {
		log.Fatal(err)
	}

	// getting a client
	client := pb.NewHelloClient(conn)

	req := &pb.HelloRequest{
		Name: "Pedro Henrique",
	}

	// calling the client's SayHello method
	res, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}
