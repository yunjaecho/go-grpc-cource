package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"../calculatorpb"
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
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Sum aUnary RPC")
	req := &calculatorpb.SumRequest{
		FirstNumber:5,
		SecondNumber:40,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalln("error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from Sum: %v", res.SumResult)
}
