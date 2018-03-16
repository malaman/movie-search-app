package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/malaman/movie-search-app/backend/http"
	"github.com/malaman/movie-search-app/backend/models"
	"github.com/malaman/movie-search-app/backend/utils"
)

func parseHttpArgs(query map[string][]string) (*models.QueryParams, error) {
	searchParam := query["s"]
	if len(searchParam) > 0 {
		return &models.QueryParams{searchParam[0]}, nil
	}
	return nil, errors.New("s query param is not found")
}

func getSearchResultItemsFromBytes(bytes *[]byte) ([]models.SearchResultItem, error) {
	searchResponse := models.SearchResponse{}
	emptyResponse := []models.SearchResultItem{}
	if err := json.Unmarshal(*bytes, &searchResponse); err != nil {
		return emptyResponse, err
	} else {
		return searchResponse.Search, nil
	}
}

func GetMovieSearchResult(query map[string][]string) ([]byte, error) {
	emptyResponse := []byte{}
	if s, err := parseHttpArgs(query); err != nil {
		return emptyResponse, err
	} else {
		url := fmt.Sprintf("%s?apikey=%s&s=%s", utils.ApiHost, utils.ApiKey, s.Search)
		if result, err := http.Get(url); err != nil {
			return emptyResponse, err
		} else {
			if searchResultItems, err := getSearchResultItemsFromBytes(&result); err != nil {
				return emptyResponse, err
			} else {
				if response, err := json.Marshal(searchResultItems); err != nil {
					return emptyResponse, err
				} else {
					return response, nil
				}
			}
		}
	}
}
