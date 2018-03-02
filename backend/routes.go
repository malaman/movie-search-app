package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
)

func parseHttpArgs(args *fasthttp.Args) {
	fmt.Println(string(args.Peek("s")))
}

func Index(ctx *fasthttp.RequestCtx) {
	parseHttpArgs(ctx.QueryArgs())
	fmt.Fprint(ctx, "Welcome!\n")
}

func StartRouter() {
	router := fasthttprouter.New()
	router.GET("/search", Index)
	log.Fatal(fasthttp.ListenAndServe(":9000", router.Handler))
}
