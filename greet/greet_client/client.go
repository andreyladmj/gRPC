package main

import (
	"andreyladmj/gRPC/greet/greetpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {

	fmt.Println("hello client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect to the server: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Stephane",
			LastName:  "Maareh",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatal("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}
