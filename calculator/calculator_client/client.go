package main

import (
	"../calculatorpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("Calculator client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	//fmt.Printf("Create client : %f", c)
	doUnary(c)

	doServerStreaming(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Sum aUnary RPC")
	req := &calculatorpb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 40,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalln("error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from Sum: %v", res.SumResult)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a CalculatorService_PrimeNumberDecompsition Streaming RPC")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 234233214569,
	}
	stream, err := c.PrimeNumberDecompsition(context.Background(), req)
	if err != io.EOF {

	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happend : %v", err)
		}
		fmt.Println(res.GetPrimeFactor())
	}

}
