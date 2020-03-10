package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/ranggakusuma/go-kit-example/utils"
)

// MovieService provide operations in movie service
type MovieService interface {
	Search(ctx context.Context, searchWord string, pagination int) (string, error)
}

type basicMovieService struct{}

// NewBasicMovieService is stateless of implementation of MovieService
func NewBasicMovieService() MovieService {
	return &basicMovieService{}
}

type responseOmdb struct {
	Response string `json:"Response"`
}

func (basicMovieService) Search(ctx context.Context, searchWord string, pagination int) (string, error) {

	client := &http.Client{}
	apiKey := os.Getenv("OMDB_KEY")
	m := responseOmdb{}

	queryURL := fmt.Sprintf("%s?apikey=%s&s=%s&page=%d", utils.OmdbURL, apiKey, url.PathEscape(searchWord), pagination)
	fmt.Println(queryURL)
	req, err := http.NewRequest("GET", queryURL, bytes.NewBuffer(nil))
	// res, err := http.Get(queryURL)
	// req.Header.Add("Content-Type", "application/json")

	if err != nil {
		// handle error
		return "", err
	}

	res, err := client.Do(req)
	if err != nil {
		return "", nil
	}
	err = json.NewDecoder(res.Body).Decode(&m)
	if err != nil {
		return "", err
	}

	//err = json.Unmarshal(json.NewDecoder(res.Body).Decode(&m))
	// resJson, err := json.Marshal(res.Body)
	// if err != nil {
	// 	return "", nil
	// }

	// json.Unmarshal(resJson, &m)

	fmt.Println(m, "response ===")
	return searchWord + "oke oke", nil
}
