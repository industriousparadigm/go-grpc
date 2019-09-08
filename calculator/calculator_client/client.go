package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/industriousparadigm/go-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Greetings, je suis une cliente calculatrice")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	// doUnary(c)

	doServerStreaming(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to perform a Unary RPC...")
	req := &calculatorpb.SumRequest{
		A: 3,
		B: 10,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}

	log.Printf("response from Sum: %v", res.SumResult)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Initiating a server stream RPC...")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		InputNumber: 327,
	}

	resStream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Server Stream RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			log.Println("Stream is done.")
			break // meaning nothing more is being streamed
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("PrimeNumberDecomposition response: %v", msg.GetResult())
	}
}
