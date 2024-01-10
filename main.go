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

	log.Printf("Server listening on :%s", c.Port)
	err = http.ListenAndServe(":" + c.Port, nil)

	if err != nil {
		log.Fatal(err)
	}
}
