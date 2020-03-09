package main

// MovieService provide operations in movie service
type MovieService interface {
	Search(searchWord string, pagination int) (string, error)
}

type movieService struct{}

func (movieService) Search(searchWord string, pagination int) (string, error) {
	return searchWord + "oke oke", nil
}
