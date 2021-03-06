package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"grpc-course/calculator/calculatorpb"

	"google.golang.org/grpc"
)

type server struct{}

func (server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber
	result := firstNumber + secondNumber
	res := &calculatorpb.SumResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Calculator Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
