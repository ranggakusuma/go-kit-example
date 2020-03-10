package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(MovieService) MovieService

type loggingMiddleware struct {
	logger log.Logger
	next   MovieService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a BugsService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next MovieService) MovieService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Search(ctx context.Context, searchWord string, pagination int) (movieList []*Movie, totalResults int, err error) {
	defer func() {
		l.logger.Log("method", "Search", "searchWord", searchWord, "pagination", pagination, "err", err)
	}()

	movieList, totalResults, err = l.next.Search(ctx, searchWord, pagination)
	return
}
