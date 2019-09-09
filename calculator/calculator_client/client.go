package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	// doServerStreaming(c)
	doClientStreaming(c)
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

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to perform a Client Streaming RPC...")

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while calling ComputeAverage: %v", err)
	}

	numbers := []int64{11, 9, 7, 5, 4}

	for _, number := range numbers {
		fmt.Printf("A request is being sent with number %v\n", number)
		stream.Send(&calculatorpb.ComputeAverageRequest{
			Number: number,
		})
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from ComputeAverage: %v", err)
	}

	fmt.Printf("Server computed an average of %v.\n", res.GetComputedAverage())
}
