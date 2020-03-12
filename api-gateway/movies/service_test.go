package movies_test

import (
	"context"
	"testing"

	"github.com/ranggakusuma/go-kit-example/api-gateway/movies"
	"github.com/ranggakusuma/go-kit-example/api-gateway/movies/mocks"
	"github.com/ranggakusuma/go-kit-example/movies/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService(t *testing.T) {
	t.Run("Test new service", func(t *testing.T) {
		movies.NewService()
	})

	t.Run("Test response service", func(t *testing.T) {
		mockService := mocks.Service{}

		mockMovies := make([]*service.Movie, 0)
		mockReturn := movies.SearchMovieResponse{
			Search:       mockMovies,
			TotalResults: 0,
		}

		mockService.On("SearchMovies", mock.Anything, mock.Anything).Return(mockReturn, nil)

		mockRequest := movies.SearchMovieRequest{
			SearchWord: "Mock",
			Pagination: 1,
		}

		svc := movies.NewService()

		res, err := svc.SearchMovies(context.Background(), mockRequest)
		if err != nil {
			panic(err)
		}

		assert.Equal(t, len(res.Search), 0, "they should be equal")

		assert.Equal(t, res.TotalResults, 0, "they should be equal")
	})

}
