package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/industriousparadigm/go-grpc/greet/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Greetings " + firstName
	response := &greetpb.GreetResponse{
		Result: result,
	}
	return response, nil
}

func main() {
	fmt.Println("Greet server is a-go!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve?! %v", err)
	}
}
