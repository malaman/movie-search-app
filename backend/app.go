package main

import (	
	"github.com/malaman/movie-search-app/backend/utils"
)

func main() {
	utils.LoadEnvVariables()
	StartRouter()
}
