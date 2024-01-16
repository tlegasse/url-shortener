package main

import (
	"log"
	"net/http"

    urlshortener "github.com/tlegasse/url-shortener/internal/urlshortener"
	util "github.com/tlegasse/url-shortener/internal/util"
	database "github.com/tlegasse/url-shortener/internal/database"
)

func main() {
	c := util.GetConfig()
	SetupUrlShortener(c)
	SetupServer(c)
}

func SetupUrlShortener(c util.Config) {
	db := database.Db
	urlshortener.Setup(c.BaseURL, c.Port, db)
}

func SetupServer(c util.Config) {
	log.Printf("Server listening on :%s", c.Port)
	err := http.ListenAndServe(":" + c.Port, nil)

	if err != nil {
		log.Fatal(err)
	}
}

