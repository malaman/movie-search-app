package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/malaman/movie-search-app/backend/http"
	"github.com/malaman/movie-search-app/backend/models"
	"github.com/malaman/movie-search-app/backend/utils"
	"github.com/valyala/fasthttp"
)

func parseHttpArgs(args *fasthttp.Args) (*models.QueryParams, error) {
	searchParam := string(args.Peek("s"))
	if len(searchParam) == 0 {
		return nil, errors.New("s query param is not found")
	}
	return &models.QueryParams{searchParam}, nil
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

func GetMovieSearchResult(ctx *fasthttp.RequestCtx) ([]byte, error) {
	emptyResponse := []byte{}
	if s, err := parseHttpArgs(ctx.QueryArgs()); err != nil {
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