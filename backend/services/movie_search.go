package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/malaman/movie-search-app/backend/http"
	"github.com/malaman/movie-search-app/backend/models"
	"github.com/malaman/movie-search-app/backend/utils"
	"strconv"
)

func parseHttpArgs(query map[string][]string) (*models.OMDBSearchQueryParams, error) {
	result := models.OMDBSearchQueryParams{"", 1}
	err := false
	for key, value := range query {
		switch {
		case key == "s" && len(value) > 0:
			result.Search = value[0]
		case key == "page" && len(value) > 0:
			if page, err := strconv.Atoi(value[0]); err != nil {
				return nil, errors.New("Unable to parse search query")
			} else {
				result.Page = page
			}
		default:
			err = true
		}
	}
	if err {
		return nil, errors.New("Unable to parse search query")
	}
	return &result, nil
}

func getSearchResultItemsFromBytes(bytes *[]byte) ([]models.OMDBSearchResultItem, error) {
	searchResponse := models.OMDBSearchResponse{}
	emptyResponse := []models.OMDBSearchResultItem{}
	if err := json.Unmarshal(*bytes, &searchResponse); err != nil {
		return emptyResponse, err
	} else {
		return searchResponse.Search, nil
	}
}

/*
func getNextLink(totalResults int, currentQuery models.OMDBSearchQueryParams) {
	switch {
	case (OMDB_ITEMS_PER_PAGE * totalResult.Page) < totalResults:

	}

}


func transformOMDBSearchResponse(result models.OMDBSearchResult) models.SearchResponse {
	emptyResponse := models.SearchResponse{}

}
*/

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
