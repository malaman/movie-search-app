package services

import (
	"testing"
)

var SearchResultMock = `
{
  "Search": [
    {
      "Title": "Lord, Give Me Patience",
      "Year": "2017",
      "imdbID": "tt5991410",
      "Type": "movie",
      "Poster": "https://images-na.ssl-images-amazon.com/images/M/MV5BOGU1ZDgxNWMtOTEyZi00MjZjLWI2Y2QtYjMxMTNmOTJlMzRmXkEyXkFqcGdeQXVyMjQ3NzUxOTM@._V1_SX300.jpg"
    },
    {
      "Title": "Dark Lord",
      "Year": "2017–",
      "imdbID": "tt5996820",
      "Type": "series",
      "Poster": "N/A"
    },
    {
      "Title": "The Lord Inquisitor: Seed of Ambition",
      "Year": "2017",
      "imdbID": "tt5172040",
      "Type": "movie",
      "Poster": "N/A"
    },
    {
      "Title": "Lord of Shadows",
      "Year": "2017",
      "imdbID": "tt5715602",
      "Type": "movie",
      "Poster": "N/A"
    }
  ],
  "totalResults": "4",
  "Response": "True"
}
`

func TestGetSearchResultItemsFromBytes(t *testing.T) {
	bytes := []byte(SearchResultMock)
	if result, err := getSearchResultItemsFromBytes(&bytes); err != nil {
		t.Error("getSearchResultItemsFromBytes is failed")
	} else {
		if len(result.Search) != 4 {
			t.Error("getSearchResultItemsFromBytes: slice len does not match")
		}
		if result.Search[0].Title != "Lord, Give Me Patience" {
			t.Error("getSearchResultItemsFromBytes: slice content does not match")
		}
	}
}
