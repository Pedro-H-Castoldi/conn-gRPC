package main

import (
	"context"
	"hello_grpc/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Server is a structure with data referring to .pb file
type Server struct {
	pb.UnimplementedHelloServer
}

// SayHello has the business logic referring to the SayHello function from the pb package
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Response: "Hello " + in.GetName()}, nil
}

func main() {
	println("Running gRPC server")

	// creating connection to server
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}

	// creating and registering a new server
	newServer := grpc.NewServer()
	pb.RegisterHelloServer(newServer, &Server{})

	// running the server
	if err := newServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
