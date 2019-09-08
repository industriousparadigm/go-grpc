package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/industriousparadigm/go-grpc/calculator/calcpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calcpb.SumRequest) (*calcpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked with %v\n", req)
	a, b := req.GetSum().GetA(), req.GetSum().GetB()
	result := a + b
	response := &calcpb.SumResponse{
		Result: result,
	}
	return response, nil
}

func main() {
	fmt.Println("Calc server is a-go!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve?! %v", err)
	}
}
