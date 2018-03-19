package models

type SearchResponse struct {
	Results      []OMDBSearchResultItem `json:"results"`
	TotalResults string                 `json:"totalResults"`
	NextLink     string                 `json:"nextLink"`
}

type SearchQueryParms struct {
	*OMDBSearchQueryParams
}
