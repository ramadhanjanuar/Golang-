package main

import (
	"context"
	"fmt"
	"log"

	"grpc-course/calculator/calculatorpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("I'm client")
	con, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer con.Close()

	calculatorService := calculatorpb.NewCalculatorServiceClient(con)
	doUnary(calculatorService)
}

func doUnary(calculatorService calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.SumRequest{
		FirstNumber:  10,
		SecondNumber: 10,
	}
	res, err := calculatorService.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call calculator RPC: %v", err)
	}
	log.Printf("Response: %v", res.Result)
}
