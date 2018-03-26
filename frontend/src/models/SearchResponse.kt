package models
external interface SearchResponseItem {
    val Title: String
    val Year: Int
    val imdbID: String
    //TODO: use enum here
    val Type: String
    val Poster: String

}
external interface SearchResponse {
    val results: Array<SearchResponseItem>
    val totalResults: String
    val nextPage: String
}