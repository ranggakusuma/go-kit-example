package movies

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type searchMovieRequest struct {
	SearchWord string `json:"searchWord"`
	Pagination int    `json:"pagination"`
}

type searchMovieResponse struct {
	V   string `json:"v"`
	Err error  `json:"error,omitempty"`
}

func makeMovieEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(searchMovieRequest)
		v, err := s.SearchMovies(ctx, req)

		return searchMovieResponse{
			V: v,
		}, nil
	}
}
