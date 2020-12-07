package grpc

import (
	"context"
	v1 "grpcCRUD/api/v1"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, v1API v1.ToDoServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	v1.RegisterToDoServiceServer(server, v1API)
	//建立一個channel接收
	c := make(chan os.Signal, 1)
	go func() {
		for range c {
			log.Println("shutting down grpc server...")
			// GracefulStop stops the gRPC server gracefully. It stops the server from
			// accepting new connections and RPCs and blocks until all the pending RPCs are finished.
			server.GracefulStop()
			<-ctx.Done()
		}
	}()
	log.Println("starting grpc server...")
	//Serve accepts incoming connections on the listener lis, creating a new ServerTransport and service goroutine for each
	return server.Serve(listen)
}
