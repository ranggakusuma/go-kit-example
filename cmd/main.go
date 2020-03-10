package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/joho/godotenv"
	"github.com/ranggakusuma/go-kit-example/endpoint"
	"github.com/ranggakusuma/go-kit-example/service"
	httphandler "github.com/ranggakusuma/go-kit-example/transport/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	logger := log.NewLogfmtLogger(os.Stderr)
	svc := service.NewBasicMovieService()

	svcEndpoints := endpoint.MakeEndpoints(svc)
	h := httphandler.NewService(svcEndpoints, logger)

	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", h))
}
