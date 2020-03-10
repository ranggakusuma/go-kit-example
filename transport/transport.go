package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/ranggakusuma/go-kit-example/endpoint"
)

// ErrBadRequest is error message for bad request
var ErrBadRequest = errors.New("Request not valid cooy")

// DecodeSearchRequest is function for decode search request
func DecodeSearchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	searchWord := r.FormValue("searchword")
	strPagination := r.FormValue("pagination")
	pagination := 0

	if strPagination == "" {
		pagination = 1
	} else {
		resInt, err := strconv.Atoi(strPagination)
		if err != nil {
			return nil, ErrBadRequest
		}
		pagination = resInt
	}

	return endpoint.SearchRequest{
		SearchWord: searchWord,
		Pagination: pagination,
	}, nil
}

// EncodeResponse is function for encode response
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
