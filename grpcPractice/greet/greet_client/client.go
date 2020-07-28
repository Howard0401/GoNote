package main

import (
	"context"
	"fmt"
	"grpcPractice/greet/greetpb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println(("Hellow I'm a client"))
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) //dont have SSL
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	// fmt.Printf("Create client: %f", c)

	// c.Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption)  from greeetpb.go

	// doUnary(c)

	// doServerStreaming(c)
	doClientStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Howard",
			LastName:  "Chen",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Printf("Start to do a Server StreamRPC...")
	//Create the requests
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Howard",
			LastName:  "Chen",
		},
	}
	//place the rpc code, get the resonse streams
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	//loop
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			//reach the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response for GreetManyTimes: %v", msg.GetResult())
	}
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC...")
	//clients' requests
	// requests := []*greetpb.LongGreetRequest{ //slice with length
	requests := []*greetpb.LongGreetRequest{ //slice with length
		// &greetpb.LongGreetReques
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Howard",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Stephane",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Lucy",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Piper",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Sam",
			},
		},
	}
	stream, err := c.LongGreet(context.Background()) //因為這次client端會一直傳資料，非單筆request，不用req引數
	if err != nil {
		log.Fatalf("error while calling LonGreet: %v", err)
	}
	//iterate over out slice and send each massage individually
	for _, req := range requests {
		fmt.Printf("Sending requests:%v\n", req)
		err := stream.Send(req)
		if err != nil {
			log.Fatalf("Get err while request: %v", err)
		}
		time.Sleep(1000 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error while receive response from LongGreet: %v", err)
	}

	fmt.Printf("LongGreet Response:%v\n", res)
}
