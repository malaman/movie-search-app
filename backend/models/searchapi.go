package models

import (
	"encoding/json"
	"reflect"

	"github.com/malaman/movie-search-app/backend/utils"
)

type SearchResponse struct {
	Results      []OMDBSearchResultItem `json:"results"`
	TotalResults string                 `json:"totalResults"`
	NextPage     string                 `json:"nextPage"`
}

type SearchQueryParms struct {
	*OMDBSearchQueryParams
}

type MovieDetails struct {
	*OMDBMovieDetails
}

func (movieDetails MovieDetails) MarshalJSON() ([]byte, error) {
	result := map[string]interface{}{}
	v := reflect.ValueOf(movieDetails.OMDBMovieDetails).Elem()
	for i := 0; i < v.NumField(); i++ {
		typeField := v.Type().Field(i)
		fieldName := utils.PascalCaseToCamelCase(typeField.Name)
		result[fieldName] = v.Field(i).Interface()
	}
	return json.Marshal(result)
}
