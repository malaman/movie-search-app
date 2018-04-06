package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/malaman/movie-search-app/backend/http"
	"github.com/malaman/movie-search-app/backend/models"
	"github.com/malaman/movie-search-app/backend/utils"
)

//TODO: add tests here
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

func getSearchResultItemsFromBytes(bytes *[]byte) (models.OMDBSearchResponse, error) {
	searchResponse := models.OMDBSearchResponse{}
	if err := json.Unmarshal(*bytes, &searchResponse); err != nil {
		return models.OMDBSearchResponse{}, err
	}
	return searchResponse, nil
}

/*
func getNextLink(totalResults int, currentQuery models.OMDBSearchQueryParams) {
	switch {
	case (OMDB_ITEMS_PER_PAGE * totalResult.Page) < totalResults:

	}

}

*/
func transformOMDBSearchResponse(result models.OMDBSearchResponse) models.SearchResponse {
	response := models.SearchResponse{
		Results:      result.Search,
		TotalResults: result.TotalResults,
		//TODO: add nextpage link calculation here
		NextPage: ""}
	return response
}

//GetMovieSearchResult prepares data for the movies search request
func GetMovieSearchResult(query map[string][]string, client http.HTTPClient) ([]byte, error) {
	emptyResponse := []byte{}
	if s, err := parseHttpArgs(query); err != nil {
		return emptyResponse, err
	} else {
		url := fmt.Sprintf("%s?apikey=%s&s=%s", utils.ApiHost, utils.ApiKey, s.Search)
		if result, err := client.Get(url); err != nil {
			return emptyResponse, err
		} else {
			if searchResponse, err := getSearchResultItemsFromBytes(&result); err != nil {
				return emptyResponse, err
			} else {
				searchResult := transformOMDBSearchResponse(searchResponse)
				if response, err := json.Marshal(searchResult); err != nil {
					return emptyResponse, err
				} else {
					return response, nil
				}
			}
		}
	}
}
