package models

type SearchResponse struct {
	Results      []OMDBSearchResultItem `json:"results"`
	TotalResults string                 `json:"totalResults"`
	NextPage     string                 `json:"nextPage"`
}

type SearchQueryParms struct {
	*OMDBSearchQueryParams
}
