package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/ranggakusuma/go-kit-example/api-gateway/movies"
)

func main() {
	var logger log.Logger

	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	var ms movies.Service
	ms = movies.NewService()
	// ms = movies.NewBasicMovieService()
	ms = movies.NewLoggingService(log.With(logger, "component", "movies"), ms)

	httpLogger := log.With(logger, "component", "http")

	http.Handle("/", movies.MakeHandler(ms, httpLogger))

	logger.Log("msg", "HTTP", "addr", ":8080")
	errs := make(chan error)

	go func() {
		logger.Log("transport", "http", "address", ":8080", "msg", "listening serv")
		errs <- http.ListenAndServe(":8080", nil)
	}()

	logger.Log("terminated", <-errs)
}
