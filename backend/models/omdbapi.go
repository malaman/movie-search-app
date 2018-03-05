package models

type SearchResponse struct {
	Search []struct {
		Title  string `json:"Title"`
		Year   string `json:"Year"`
		ImdbID string `json:"imdbID"`
		Type   string `json:"Type"`
		Poster string `json:"Poster"`
	} `json:"Search"`
	TotalResults string `json:"totalResults"`
	Response     string `json:"Response"`
}

type QueryParams struct {
	Search string 
}
