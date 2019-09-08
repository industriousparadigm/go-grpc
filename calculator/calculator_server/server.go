package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/industriousparadigm/go-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked with %v\n", req)
	a, b := req.GetA(), req.GetB()
	sum := a + b
	res := &calculatorpb.SumResponse{
		SumResult: sum,
	}
	return res, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("Prime Number Decomposition function was invoked with %v\n", req)

	n := req.GetInputNumber()
	var divider int32 = 2

	for n > 1 {
		if n%divider == 0 {
			res := &calculatorpb.PrimeNumberDecompositionResponse{
				Result: divider,
			}
			stream.Send(res)
			n = n / divider
		} else {
			divider = divider + 1
		}
		fmt.Printf("current values: n=%v ; divider=%v\n", n, divider)
	}

	return nil
}

func main() {
	fmt.Println("Calc server is a-go!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve?! %v", err)
	}
}
