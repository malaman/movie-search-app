package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	CustomHTTP "github.com/malaman/movie-search-app/backend/http"
	"github.com/malaman/movie-search-app/backend/services"
	"github.com/rs/cors"
)

var client = &CustomHTTP.Client{}

func search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	w.Header().Set("Content-Type", "application/json")
	if result, err := services.GetMovieSearchResult(query, client); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("{\"error\": \"No result\"}"))
	} else {
		w.Write(result)
	}
}

func details(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	if result, err := services.GetMovieDetails(params["imdbID"], client); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("{\"error\": \"No result\"}"))
	} else {
		w.Write(result)
	}
}

//StartRouter - starts mux router
func StartRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/search", search)
	router.HandleFunc("/movie/{imdbID}", details)

	handler := cors.Default().Handler(router)
	log.Println("Starting a server on port :9000")
	log.Fatal(http.ListenAndServe(":9000", handler))
}
