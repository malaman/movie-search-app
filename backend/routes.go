package main

import (	
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
	"github.com/malaman/movie-search-app/backend/services"		
)

func Search(ctx *fasthttp.RequestCtx) {
	if result, err := services.GetMovieSearchResult(ctx); err != nil {
		log.Println(err)
		ctx.Error(fasthttp.StatusMessage(fasthttp.StatusNotFound), fasthttp.StatusNotFound)		
		fmt.Fprintf(ctx, "Error")
	} else {
		fmt.Fprintf(ctx, result)		
	}	
}

func StartRouter() {
	router := fasthttprouter.New()
	router.GET("/search", Search)
	log.Fatal(fasthttp.ListenAndServe(":9000", router.Handler))
}
