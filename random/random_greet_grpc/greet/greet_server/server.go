package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"grpc-course/greet/greetpb"
)

type server struct{}

func (server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	firstName := req.GetGreeting().GetFirstName()
	result := "Halo" + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello")
	lis, err := net.Listen("tcp", "0.0.0.0:50010")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := NewServerWithAuthInterceptor()
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
