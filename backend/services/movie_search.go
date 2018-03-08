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

func ParseHttpArgs(args *fasthttp.Args) (*models.QueryParams, error) {
	searchParam := string(args.Peek("s"))
	if len(searchParam) == 0 {
		return nil, errors.New("s query param is not found")
	}
	return &models.QueryParams{searchParam}, nil
}

func GetMovieSearchResult(ctx *fasthttp.RequestCtx) (*[]byte, error) {
	emptyResponse := []byte{}
	if s, err := ParseHttpArgs(ctx.QueryArgs()); err != nil {
		return &emptyResponse, err
	} else {
		url := fmt.Sprintf("%s?apikey=%s&s=%s", utils.ApiHost, utils.ApiKey, s.Search)
		if result, err := http.Get(url); err != nil {
			return &emptyResponse, err
		} else {
			searchResponse := models.SearchResponse{}
			if err := json.Unmarshal(result, &searchResponse); err != nil {
				return &emptyResponse, err
			} else {
				if response, err := json.Marshal(searchResponse.Search); err != nil {
					return &emptyResponse, err
				} else {
					return &response, nil
				}
			}
		}
	}
}
