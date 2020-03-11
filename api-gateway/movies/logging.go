package movies

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) SearchMovies(ctx context.Context, req searchMovieRequest) (v string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "search",
			"searchWord", req.SearchWord,
			"pagination", req.Pagination,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.Service.SearchMovies(ctx, req)
}
