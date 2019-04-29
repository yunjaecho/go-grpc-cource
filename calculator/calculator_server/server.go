package main

import (
	"../calculatorpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (*server) Sum(ctx context.Context, request *calculatorpb.SumRequest) (*calculatorpb.SumRespnse, error) {
	fmt.Printf("Received Sum RPC: %v", request)
	firstNumber := request.FirstNumber
	secondNumber := request.SecondNumber
	sum := firstNumber + secondNumber
	res := &calculatorpb.SumRespnse{
		SumResult: sum,
	}
	return res, nil
}

func (*server) PrimeNumberDecompsition(request *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompsitionServer) error {
	fmt.Printf("Received PrimeNumberDecomposition RPC: %v\n", request)
	number := request.GetNumber()
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			fmt.Printf("Divisor has increased to %v", divisor)
		}
	}
	return nil
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
