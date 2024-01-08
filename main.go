package main

import (
	"log"
	"net/http"

    urlshortener "github.com/tlegasse/url-shortener/internal/urlshortener"
    database "github.com/tlegasse/url-shortener/internal/database"
)

func main() {
	defer database.DbInstance.Close()

	u := urlshortener.ShortenerInstance

	http.HandleFunc("/shorten", u.Shorten)
	http.HandleFunc("/", u.Redirect)

	log.Println("Server listening on :8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
