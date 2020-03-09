package main

import (
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {
	svc := movieService{}
	r := mux.NewRouter()

	searchHandler := httptransport.NewServer(
		MakeSearchEndpoint(svc),
		decodeSearchRequest,
		encodeResponse,
	)

	r.Methods("GET").Path("/search").Handler(searchHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
