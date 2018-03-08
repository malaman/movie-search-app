package models

type SearchResultItem struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type SearchResponse struct {
	Search       []SearchResultItem `json:"Search"`
	TotalResults string             `json:"totalResults"`
	Response     string             `json:"Response"`
}

type QueryParams struct {
	Search string
}
