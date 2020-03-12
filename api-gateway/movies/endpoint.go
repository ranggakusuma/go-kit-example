package movies

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/ranggakusuma/go-kit-example/movies/service"
)

type SearchMovieRequest struct {
	SearchWord string `json:"searchWord"`
	Pagination int    `json:"pagination"`
}

type SearchMovieResponse struct {
	Search       []*service.Movie `json:"search"`
	TotalResults int              `json:"totalResults"`
	Err          error            `json:"error,omitempty"`
}

func makeMovieEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SearchMovieRequest)
		v, err := s.SearchMovies(ctx, req)

		if err != nil {
			return SearchMovieResponse{
				Err: err,
			}, nil
		}

		return v, nil
	}
}
