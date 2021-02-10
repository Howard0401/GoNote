package main

import (
	"context"
	"fmt"
	pb "gRPCTimeout/pb"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
}

func main() {
	fmt.Println("hi, welcome to server..")
	ctx := context.Background()
	listen, _ := net.Listen("tcp", "0.0.0.0:50051")
	grpcServer := grpc.NewServer()

	clientDeadline := time.Now().Add(time.Duration(90*time.Second) * time.Millisecond)
	ctx, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()

	// Create a cmux.
	// m := cmux.New(listen)
	// grpcL := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	// s.TestGreet(ctx, *pb.GreetRequest)
	// greeting := pb.RegisterHelloServer() //.RegisterHelloServer(s, &server{})
	s := &Server{}
	pb.RegisterHelloServer(grpcServer, s)
	c := make(chan os.Signal, 1)
	// time.Sleep(70 * time.Second)
	go func() {
		for range c {
			grpcServer.GracefulStop()
			<-ctx.Done()
		}
	}()
	grpcServer.Serve(listen)
	// go grpcServer.Serve(grpcL)
	// m.Serve()
}

func (s *Server) TestGreet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	time.Sleep(10 * time.Second)
	select {
	default:
	case <-ctx.Done():
		log.Panicf("err: %v", ctx.Err())
		return nil, ctx.Err()
	}
	req.Say = fmt.Sprintf("Received: %v", req.Say)
	log.Printf("Receive:%v", req.Say)
	return &pb.GreetResponse{Receive: req.Say}, nil
}
