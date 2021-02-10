package main

import (
	"context"
	"gRPCTimeout/pb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Conn Err: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloClient(conn)
	res, err := c.TestGreet(ctx, &pb.GreetRequest{Say: "golang grpc000"})
	select {
	case <-ctx.Done():
		log.Panicf("err: %v", ctx.Err())
	default:
		log.Println("ctx works well")
	}
	if err != nil {
		log.Fatalf("c.TestGreet Error:%v", err)
	}
	log.Printf("res = %v", res)
}
