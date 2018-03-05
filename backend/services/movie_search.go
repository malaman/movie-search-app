package services

import (
	"fmt"
	"errors"		
	"github.com/valyala/fasthttp"	
    "github.com/malaman/movie-search-app/backend/models"	
	"github.com/malaman/movie-search-app/backend/utils"	
	"github.com/malaman/movie-search-app/backend/http"	
)

func ParseHttpArgs(args *fasthttp.Args) (*models.QueryParams, error) {
	searchParam := string(args.Peek("s"))
    if (len(searchParam)  == 0 ) {
		return nil, errors.New("s query param is not found")		
	}	
	return &models.QueryParams{searchParam}, nil
}

func GetMovieSearchResult(ctx *fasthttp.RequestCtx) (string, error) {
	if s, err := ParseHttpArgs(ctx.QueryArgs()); err != nil {		
		return "", err		
	} else {		
		url := fmt.Sprintf("%s?apikey=%s&s=%s", utils.ApiHost, utils.ApiKey, s.Search)
		if result, err := http.Get(url); err != nil {
			return "", err
		} else {
			return result, nil
		}
	}
}
