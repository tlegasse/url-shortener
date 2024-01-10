package main

import (
	"log"
	"net/http"

    urlshortener "github.com/tlegasse/url-shortener/internal/urlshortener"
	util "github.com/tlegasse/url-shortener/internal/util"
)

func main() {
	c := GetConfig()
	SetupUrlShortener(c)
	SetupServer(c)
}

func SetupUrlShortener(c util.Config) {
	urlshortener.Setup(c.BaseURL, c.Port)
}

func SetupServer(c util.Config) {
	log.Printf("Server listening on :%s", c.Port)
	err := http.ListenAndServe(":" + c.Port, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func GetConfig() util.Config {
	c, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	return c
}
