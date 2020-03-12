package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/joho/godotenv"
	"github.com/ranggakusuma/go-kit-example/movies/endpoint"
	"github.com/ranggakusuma/go-kit-example/movies/service"

	grpchandler "github.com/ranggakusuma/go-kit-example/movies/transport/grpc"
	"github.com/ranggakusuma/go-kit-example/movies/transport/grpc/pb"
	grpc "google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file not found")
	}

	logger := log.NewLogfmtLogger(os.Stderr)
	svc := service.NewBasicMovieService()
	ctx := context.Background()

	errChan := make(chan error)

	svcEndpoints := endpoint.MakeEndpoints(svc)

	go func() {
		listener, err := net.Listen("tcp", ":8081")

		if err != nil {
			errChan <- err
			return
		}

		h := grpchandler.NewGRPCServer(ctx, svcEndpoints, logger)

		gRPCServer := grpc.NewServer()
		pb.RegisterMovieServer(gRPCServer, h)
		errChan <- gRPCServer.Serve(listener)
	}()

	logger.Log("msg", "gRPC", "addr", ":8081")
	fmt.Println(<-errChan)
}
