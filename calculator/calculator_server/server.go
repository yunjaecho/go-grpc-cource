package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"../calculatorpb"
)

type server struct{}

func (*server) Sum(ctx context.Context, request *calculatorpb.SumRequest) (*calculatorpb.SumRespnse, error) {
	fmt.Printf("Received Sum RPC: %v", request)
	firstNumber := request.FirstNumber
	secondNumber := request.SecondNumber
	sum := firstNumber + secondNumber
	res := &calculatorpb.SumRespnse{
		SumResult:sum,
	}
	return res, nil
}

func main() {
	fmt.Println("Calculator Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}