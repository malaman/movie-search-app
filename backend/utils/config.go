package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var ApiHost string
var ApiKey string

func LoadEnvVariables() {	
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
    ApiHost = os.Getenv("API_HOST")
    ApiKey = os.Getenv("API_KEY")
}
