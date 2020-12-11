package grpc

import (
	"context"
	v1 "grpcCRUD/api/v1"
	middleware "grpcCRUD/pkg/grpc/middleware"
	logger "grpcCRUD/pkg/logger"
	"net"
	"os"

	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, v1API v1.ToDoServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	opts := []grpc.ServerOption{}
	//add middleware which contains gRPC server statup options 藉由中間層將logger package設定格式輸入gRPCserver中
	opts = middleware.AddLogging(logger.Log, opts)

	//register server 建立伺服器
	server := grpc.NewServer(opts...)
	v1.RegisterToDoServiceServer(server, v1API)

	// create a channel to receive msgs 建立一個channel接收訊息
	c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// log.Println("shutting down grpc server...")
			logger.Log.Warn("shutting down gRPC server...")
			// GracefulStop stops the gRPC server gracefully. It stops the server from
			// accepting new connections and RPCs and blocks until all the pending RPCs are finished.
			server.GracefulStop()
			<-ctx.Done()
		}
	}()
	// log.Println("starting grpc server...")
	logger.Log.Info("starting gRPC server")
	//Serve accepts incoming connections on the listener lis, creating a new ServerTransport and service goroutine for each
	return server.Serve(listen)
}
