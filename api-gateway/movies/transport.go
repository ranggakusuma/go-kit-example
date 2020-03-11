package movies

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// MakeHandler is create new Handler for movies service
func MakeHandler(svc Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{}

	movieHandler := kithttp.NewServer(
		makeMovieEndpoint(svc),
		decodeSearchMoviesRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()
	r.Handle("/search", movieHandler).Methods("GET")

	return r
}

func decodeSearchMoviesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	searchWord := r.FormValue("searchword")
	strPagination := r.FormValue("pagination")
	pagination := 0

	if strPagination == "" {
		pagination = 1
	} else {
		resInt, err := strconv.Atoi(strPagination)
		if err != nil {
			return searchMovieRequest{
				SearchWord: searchWord,
				Pagination: 1,
			}, nil
		}
		pagination = resInt
	}

	return searchMovieRequest{
		SearchWord: searchWord,
		Pagination: pagination,
	}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
