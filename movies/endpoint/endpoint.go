package endpoint

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/ranggakusuma/go-kit-example/service"
)

// SearchRequest represent request for Search method
type SearchRequest struct {
	SearchWord string
	Pagination int
}

// SearchResponse represent response for Search method
type SearchResponse struct {
	Search       []*service.Movie `json:"search"`
	TotalResults int              `json:"totalResults"`
	Err          string           `json:"error,omitempty"`
}

// ErrRequestNotFound is error message for request not found
var ErrRequestNotFound = errors.New("Request not found")

// Endpoints wrapper
type Endpoints struct {
	SearchEndpoint endpoint.Endpoint
}

// MakeEndpoints initializes all Go kit endpoints
func MakeEndpoints(s service.MovieService) Endpoints {
	return Endpoints{
		SearchEndpoint: MakeSearchEndpoint(s),
	}
}

// MakeSearchEndpoint is endpoint for Search method
func MakeSearchEndpoint(svc service.MovieService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SearchRequest)
		results, totalResults, err := svc.Search(ctx, req.SearchWord, req.Pagination)
		if err != nil {
			return SearchResponse{Err: err.Error()}, nil
		}
		return SearchResponse{
			Search:       results,
			TotalResults: totalResults,
		}, nil
	}
}
