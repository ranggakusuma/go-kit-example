package movies

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/ranggakusuma/go-kit-example/movies/service"
)

type searchMovieRequest struct {
	SearchWord string `json:"searchWord"`
	Pagination int    `json:"pagination"`
}

type searchMovieResponse struct {
	Search       []*service.Movie `json:"search"`
	TotalResults int              `json:"totalResults"`
	Err          error            `json:"error,omitempty"`
}

func makeMovieEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(searchMovieRequest)
		v, err := s.SearchMovies(ctx, req)

		if err != nil {
			return searchMovieResponse{
				Err: err,
			}, nil
		}

		return v, nil
	}
}
