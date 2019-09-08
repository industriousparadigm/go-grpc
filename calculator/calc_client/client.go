package main

import (
	"context"
	"fmt"
	"log"

	"github.com/industriousparadigm/go-grpc/calculator/calcpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Greetings, je suis un client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := calcpb.NewCalculatorServiceClient(cc)

	doUnary(c)
}

func doUnary(c calcpb.CalculatorServiceClient) {
	fmt.Println("Starting to perform a Unary RPC...")
	req := &calcpb.SumRequest{
		Sum: &calcpb.Sum{
			A: 3,
			B: 10,
		},
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}

	log.Printf("response from Sum: %v", res.Result)
}
