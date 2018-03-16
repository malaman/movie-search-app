package main

import (
	"github.com/malaman/movie-search-app/backend/services"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if result, err := services.GetMovieSearchResult(query); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"error\": \"No result\"}"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

func StartRouter() {
	mux := http.NewServeMux()
	mux.HandleFunc("/search", Search)
	handler := cors.Default().Handler(mux)
	log.Println("Starting a server on port :9000")
	log.Fatal(http.ListenAndServe(":9000", handler))
}
