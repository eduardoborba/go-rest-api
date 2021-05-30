package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	a := App{}
	gopath := os.Getenv("GOPATH")
	err := godotenv.Load(gopath + "/src/github.com/eduardoborba/go-rest-api/.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8010")
}
