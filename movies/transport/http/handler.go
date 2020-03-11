package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ranggakusuma/go-kit-example/movies/endpoint"
	"github.com/ranggakusuma/go-kit-example/movies/transport"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

// NewService wires Go kit endpoints to the HTTP transport.
func NewService(svcEndpoints endpoint.Endpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	options := []httptransport.ServerOption{httptransport.ServerErrorLogger(logger)}

	r.Methods("GET").Path("/search").Handler(httptransport.NewServer(
		svcEndpoints.SearchEndpoint,
		transport.DecodeSearchRequest,
		transport.EncodeResponse,
		options...,
	))

	return r
}
