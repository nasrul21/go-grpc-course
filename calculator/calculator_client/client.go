package main

import (
	"context"
	"fmt"
	"go-grpc-course/calculator/calculatorpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting Calculator Client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := calculatorpb.NewCalculatorServiceClient(conn)

	req := &calculatorpb.SumRequest{
		Sum: &calculatorpb.Sum{
			FirstNumber:  3,
			SecondNumber: 10,
		},
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}

	log.Printf("Response from Sum Calculator: %v", res.Result)
}
