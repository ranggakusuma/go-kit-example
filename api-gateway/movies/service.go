package movies

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/ranggakusuma/go-kit-example/movies/service"
	"github.com/ranggakusuma/go-kit-example/movies/transport/grpc/pb"
	"google.golang.org/grpc"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// Service is the interface that provides list of movies service.
type Service interface {
	SearchMovies(ctx context.Context, request SearchMovieRequest) (SearchMovieResponse, error)
}

type movieService struct {
	movieServiceClient pb.MovieClient
}

// NewService is creating new service for movies
func NewService() Service {
	grpcAddr := os.Getenv("MOVIE_GRPC_ADDRESS")
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(10*time.Second))

	if err != nil {
		log.Printf("Unable to connect to movie service: %s", err.Error())
		return new(movieService)
	}

	return &movieService{
		movieServiceClient: pb.NewMovieClient(conn),
	}

}

func (m movieService) SearchMovies(ctx context.Context, request SearchMovieRequest) (SearchMovieResponse, error) {
	reply, err := m.movieServiceClient.Search(context.Background(), &pb.SearchMovieRequest{
		SearchWord: request.SearchWord,
		Pagination: int64(request.Pagination),
	})

	if err != nil {
		return SearchMovieResponse{
			Err: err,
		}, nil
	}

	searchResults := make([]*service.Movie, 0)
	for _, val := range reply.GetList() {
		searchResults = append(searchResults, &service.Movie{
			Title:  val.GetTitle(),
			Year:   val.GetYear(),
			ImdbID: val.GetImdbID(),
			Type:   val.GetType(),
			Poster: val.GetPoster(),
		})
	}

	return SearchMovieResponse{
		Search:       searchResults,
		TotalResults: int(reply.GetTotalResult()),
	}, nil

}
