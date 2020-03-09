package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// ErrBadRequest is error message for bad request
var ErrBadRequest = errors.New("Request not valid cooy")

func decodeSearchRequest(_ context.Context, r *http.Request) (interface{}, error) {
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

	return SearchRequest{
		SearchWord: searchWord,
		Pagination: pagination,
	}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
