package main

import (
	"log"
	"net/http"

    urlshortener "github.com/tlegasse/url-shortener/internal/urlshortener"
	util "github.com/tlegasse/url-shortener/internal/util"
)

func main() {
	c, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	urlshortener.Shortener.Setup(c.BaseURL, c.Port)

	http.HandleFunc("/shorten", urlshortener.Shortener.Shorten)
	http.HandleFunc("/", urlshortener.Shortener.Redirect)

	log.Printf("Server listening on :%s", c.Port)
	err = http.ListenAndServe(":" + c.Port, nil)

	if err != nil {
		log.Fatal(err)
	}
}
