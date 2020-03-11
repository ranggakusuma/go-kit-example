package main

import (
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
		panic(err)
	}

	logger := log.NewLogfmtLogger(os.Stderr)
	svc := service.NewBasicMovieService()

	errChan := make(chan error)

	//	svcEndpoints := endpoint.MakeEndpoints(svc)
	// h := httphandler.NewService(svcEndpoints, logger)
	//	h := grpchandler.NewService(svcEndpoints, logger)

	//logger.Log("msg", "HTTP", "addr", ":8080")
	// logger.Log("err", http.ListenAndServe(":8080", grpc))

	listener, err := net.Listen("tcp", "8081")
	logger.Log("msg", "gRPC", "addr", ":8081")

	if err != nil {
		errChan <- err
		return
	}
	svcEndpoints := endpoint.MakeEndpoints(svc)
	h := grpchandler.NewGRPCServer(svcEndpoints, logger)

	gRPCServer := grpc.NewServer()
	pb.RegisterMovieServer(gRPCServer, h)
	errChan <- gRPCServer.Serve(listener)

	fmt.Println(<-errChan)
}
