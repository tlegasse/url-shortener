package main

import (
	"log"
	"net/http"

    urlshortener "github.com/tlegasse/url-shortener/internal/urlshortener"
	util "github.com/tlegasse/url-shortener/internal/util"
)

func main() {
	c := util.GetConfig()
	SetupUrlShortener(c)
	SetupServer(c)
}

func SetupUrlShortener(c util.Config) {
	urlshortener.Setup(c.BaseURL, c.Port)
}

func SetupServer(c util.Config) {
	mux := http.NewServeMux()

	mux.HandleFunc("/shorten", urlshortener.Shortener.Shorten)
	mux.HandleFunc("/", urlshortener.Shortener.Redirect)

	log.Printf("Server listening on :%s", c.Port)
	err := http.ListenAndServe(":" + c.Port, mux)

	if err != nil {
		log.Fatal(err)
	}
}

