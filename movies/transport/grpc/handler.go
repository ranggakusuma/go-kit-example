package grpc

import (
	"context"

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
func NewGRPCServer(_ context.Context, endpoints endpoint.Endpoints, logger log.Logger) pb.MovieServer {
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

// decodeSearchRequest is a transport/grpc.DecodeRequestFunc that converts a
func decodeSearchRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.SearchMovieRequest)
	return endpoint.SearchRequest{
		SearchWord: req.SearchWord,
		Pagination: int(req.Pagination),
	}, nil
}

// encodeSearchResponse is a transport/grpc.EncodeResponseFunc that converts
func encodeSearchResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.SearchResponse)

	list := make([]*pb.MovieAttr, 0)

	for _, v := range resp.Search {
		list = append(list, &pb.MovieAttr{
			Title:  v.Title,
			Year:   v.Year,
			ImdbID: v.ImdbID,
			Poster: v.Poster,
			Type:   v.Type,
		})
	}

	return &pb.SearchMovieResponse{
		List:        list,
		TotalResult: int64(resp.TotalResults),
		Err:         resp.Err,
	}, nil
}
