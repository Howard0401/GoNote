package rest

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	v1 "grpcCRUD/api/v1"

	middleware "grpcCRUD/pkg/rest/middleware"

	logger "grpcCRUD/pkg/logger"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime" //注意是v2 v2!!!

	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, grpcPort, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := v1.RegisterToDoServiceHandlerFromEndpoint(ctx, mux, "127.0.0.1:"+grpcPort, opts); err != nil {
		// log.Fatalf("failed to start http gateway: %v", err)
		log.Fatalf("failed to start http gateway: %v", err)
	}

	srv := &http.Server{
		Addr: ":" + httpPort,
		// Handler: mux,
		Handler: middleware.AddRequestID(middleware.AddLogger(logger.Log, mux)),
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	//open channel and add restrictions 建立管道並設定超時關閉條件
	go func() {
		for range c {

		}
		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	// log.Println("Start REST Service")
	log.Println("starting HTTP/REST gateway")
	return srv.ListenAndServe()
}
