package main

import (
	"../greetpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("Hello I'am a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	//fmt.Printf("Create client : %f", c)
	doUnary(c)

	doServerStreaming(c)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do streaming RPC")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Stephane",
			LastName:  "Maarek",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Resonse from GreetManyTimes: %v", msg.GetResult())
	}

}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do aUnary RPC")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Stephane",
			LastName:  "Maarek",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalln("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}
