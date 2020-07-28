package main

import (
	"context"
	"fmt"
	"grpcPractice/greet/greetpb"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

//Implement type GreetServiceServer interface in greet.pb.go
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet Function")
	firstName := req.GetGreeting().GetFirstName()
	result := "Hellow " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

//Implement type GreetServiceServer interface in greet.pb.go
func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes Function\n")
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hellow" + firstName + "number" + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		err := stream.Send(res)
		if err != nil {
			log.Fatalf("GreetManyTimes gets %v", err)
		}
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

//stream 後面接的是 Service...Server
//Implement type GreetServiceServer interface in greet.pb.go
func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Printf("LongGreet Function was invoked a streaming request\n")
	result := ""
	for {
		req, err := stream.Recv() //一直接受
		if err == io.EOF {
			//finished reading
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			}) //結束 return 是因為 SendAndClose 本身就回傳 error 如果真的有錯就會直接報錯
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		result += "Hellow " + firstName + "! "
	}
}

func main() {
	fmt.Println("hi")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Fail to Listem:%v", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Fail to serve:%v", err)
	}

}
