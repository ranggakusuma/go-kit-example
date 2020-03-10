package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/ranggakusuma/go-kit-example/utils"
)

// MovieService provide operations in movie service
type MovieService interface {
	Search(ctx context.Context, searchWord string, pagination int) ([]*Movie, int, error)
}

type basicMovieService struct{}

// Movie represent response movie from omdb
type Movie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

// NewBasicMovieService is stateless of implementation of MovieService
func NewBasicMovieService() MovieService {
	return &basicMovieService{}
}

func (basicMovieService) Search(ctx context.Context, searchWord string, pagination int) ([]*Movie, int, error) {

	//client := &http.Client{}
	apiKey := os.Getenv("OMDB_KEY")
	m := map[string]interface{}{}
	result := []*Movie{}

	queryURL := fmt.Sprintf("%s?apikey=%s&s=%s&page=%d", utils.OmdbURL, apiKey, url.PathEscape(searchWord), pagination)
	//req, err := http.NewRequest("GET", queryURL, bytes.NewBuffer(nil))

	// if err != nil {
	//	return nil, 0, err
	// }

	//res, err := client.Do(req)
	res, err := http.Get(queryURL)
	if err != nil {
		return nil, 0, nil
	}
	err = json.NewDecoder(res.Body).Decode(&m)
	if err != nil {
		return nil, 0, err
	}

	if m["Response"].(string) == "False" {
		return nil, 0, errors.New(m["Error"].(string))
	}

	for _, v := range m["Search"].([]interface{}) {
		mov := new(Movie)
		b, _ := json.Marshal(v)
		_ = json.Unmarshal(b, &mov)
		result = append(result, mov)
	}

	totalPage, _ := strconv.Atoi(m["totalResults"].(string))
	return result, totalPage, nil
}
