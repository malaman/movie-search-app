package services

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/malaman/movie-search-app/backend/http"
	"github.com/malaman/movie-search-app/backend/models"
	"github.com/malaman/movie-search-app/backend/utils"
)

func getMovieDetailsFromBytes(bytes *[]byte) (*models.OMDBMovieDetails, error) {
	movieDetails := models.OMDBMovieDetails{}
	if err := json.Unmarshal(*bytes, &movieDetails); err != nil {
		return &movieDetails, err
	}
	return &movieDetails, nil
}

//GetMovieDetails prepare data for movie details route
func GetMovieDetails(imdbID string, client http.HTTPClient) ([]byte, error) {
	emptyResponse := []byte{}
	defaultError := errors.New("Unable to provide data")
	if len(imdbID) == 0 {
		return nil, defaultError
	}
	url := fmt.Sprintf("%s?apikey=%s&i=%s", utils.ApiHost, utils.ApiKey, imdbID)
	if result, err := client.Get(url); err != nil {
		return emptyResponse, err
	} else {
		if movieDetails, err := getMovieDetailsFromBytes(&result); err != nil || movieDetails.Response != "True" {
			return emptyResponse, err
		} else {
			if response, err := json.Marshal(movieDetails); err != nil {
				return emptyResponse, err
			} else {
				return response, nil
			}
		}
	}
}
