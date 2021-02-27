package main

import (
	"context"
	"fmt"
	"go-grpc-course/calculator/calculatorpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Calculator function was invoked with %v \n", req)
	firstNumber := req.GetSum().GetFirstNumber()
	secondNumber := req.GetSum().GetSecondNumber()
	result := firstNumber + secondNumber
	res := &calculatorpb.SumResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	fmt.Println("Starting Calculator Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen : %v \n", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}
