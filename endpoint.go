package main

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

// SearchRequest represent request for Search method
type SearchRequest struct {
	SearchWord string
	Pagination int
}

// SearchResponse represent response for Search method
type SearchResponse struct {
	Title  string `json:"title"`
	Year   string `json:"year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"type"`
	Poster string `json:"poster"`
}

// ErrRequestNotFound is error message for request not found
var ErrRequestNotFound = errors.New("Request not found")

// Endpoints wrapper
type Endpoints struct {
	LoremEndpoint endpoint.Endpoint
}

// MakeSearchEndpoint is endpoint for Search method
func MakeSearchEndpoint(svc MovieService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		return SearchResponse{}, nil
	}
}
