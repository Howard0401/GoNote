package main

import (
	"context"
	"fmt"
	"grpcPractice/calculator/calculatorpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

//Implement type CalculatorServiceClient interface in calculator.pb.go
func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Receive Sum RPC: %v", req)
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber
	sum := firstNumber + secondNumber
	res := &calculatorpb.SumResponse{
		SumRequest: sum, //why in this
	}
	return res, nil
}

//Implement type CalculatorServiceClient interface in calculator.pb.go
func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecopositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {

	fmt.Printf("Received PrimeNumberDecomposiotion RPC %v\n", req)
	number := req.GetNumber()
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			//注意這邊要加檢查回傳值，教程的時候還沒有要求這個回傳值
			err := stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			})
			if err != nil {
				log.Fatalf("res failed")
			}
			number = number / divisor
		} else {
			divisor++
			fmt.Printf("Divisor has increased to %v\n", divisor)
		}
	}
	return nil
}

func main() {
	fmt.Println("Calculator Server")
	lis, err := net.Listen("tcp", "0.0.0.0:50052") //注意這邊要跟對接的Client端相同
	if err != nil {
		log.Fatalf("Fail to Listen: %v", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
