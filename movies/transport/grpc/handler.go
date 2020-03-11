package grpc

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/ranggakusuma/go-kit-example/movies/endpoint"
	"github.com/ranggakusuma/go-kit-example/movies/transport/grpc/pb"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	search grpc.Handler
}

// NewGRPCServer func
func NewGRPCServer(endpoints endpoint.Endpoints, logger log.Logger) pb.MovieServer {
	return &grpcServer{search: makeSearchHandler(endpoints, []grpc.ServerOption{grpc.ServerErrorLogger(logger)})}
}

func (g *grpcServer) Search(ctx context.Context, req *pb.SearchMovieRequest) (*pb.SearchMovieResponse, error) {
	_, rep, err := g.search.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SearchMovieResponse), nil
}

// NewService wires Go kit endpoints to the gRPC transport.
// func NewService(svcEndpoints endpoint.Endpoints, logger log.Logger) grpc.Handler {
//	return makeSearchHandler(svcEndpoints, []grpc.ServerOption{grpc.ServerErrorLogger(logger)})
// }

func makeSearchHandler(endpoint endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoint.SearchEndpoint, decodeSearchRequest, encodeSearchResponse, options...)
}

// decodeSendEmailResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain SendEmail request.
// TODO implement the decoder
func decodeSearchRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Notificator' Decoder is not impelemented")
}

// encodeSendEmailResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeSearchResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Notificator' Encoder is not impelemented")
}
