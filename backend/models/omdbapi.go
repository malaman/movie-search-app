package models

// omdbapi does not allow to specify custom limit per page
const OMDB_ITEMS_PER_PAGE = 10

type OMDBSearchResultItem struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type OMDBSearchResponse struct {
	Search       []OMDBSearchResultItem `json:"Search"`
	TotalResults string                 `json:"totalResults"`
	Response     string                 `json:"Response"`
}

type OMDBSearchQueryParams struct {
	Search string
	Page   int
}
