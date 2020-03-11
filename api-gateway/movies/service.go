package movies

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ranggakusuma/go-kit-example/movies/transport/grpc/pb"
	"google.golang.org/grpc"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// Service is the interface that provides list of movies service.
type Service interface {
	SearchMovies(ctx context.Context, request searchMovieRequest) (string, error)
}

type movieService struct {
	movieServiceClient pb.MovieClient
}

// NewService is creating new service for movies
func NewService() Service {
	// 	return &movieService{}
	grpcAddr := os.Getenv("MOVIE_GRPC_ADDRESS")
	fmt.Println(grpcAddr, " ===== oy")
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(10*time.Second))

	if err != nil {
		log.Printf("Unable to connect to movie service: %s", err.Error())
		return new(movieService)
	}
	//	defer conn.Close()

	// reply, err := pb.NewMovieClient(conn).Search(context.Background(), &pb.SearchMovieRequest{"", 1}, grpc.DialOption{grpc.WithTimeout(5 * time.Second)})

	return &movieService{
		movieServiceClient: pb.NewMovieClient(conn),
	}

}

// NewBasicMovieService is creating new service grpc
// func NewBasicMovieService() Service {
// 	grpcAddr := os.Getenv("MOVIE_GRPC_ADDRESS")
// 	fmt.Println(grpcAddr, " ===== oy")
// 	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(2*time.Second))
//
// 	if err != nil {
// 		log.Printf("Unable to connect to movie service: %s", err.Error())
// 		return new(movieService)
// 	}
// 	defer conn.Close()
//
// 	return &movieService{
// 		movieServiceClient: pb.NewMovieClient(conn),
// 	}
// }

func (m movieService) SearchMovies(ctx context.Context, request searchMovieRequest) (string, error) {
	reply, err := m.movieServiceClient.Search(context.Background(), &pb.SearchMovieRequest{
		SearchWord: request.SearchWord,
		Pagination: int64(request.Pagination),
	})

	fmt.Println(reply, "=== reply")
	fmt.Println(err, "=====")

	if err != nil {
		return "", nil
	}

	fmt.Println(reply, "hasiiil")
	return "hallo hallo bandung", nil
}
