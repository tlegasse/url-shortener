package main

import (
	"log"
	"net/http"
    "github.com/tlegasse/url-shortener/database"
    "github.com/tlegasse/url-shortener/urlshortener"
)

func main() {
	db, err := database.Setup()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	shortener := urlshortener.New(db)

	http.HandleFunc("/shorten", shortener.Shorten)
	http.HandleFunc("/", shortener.Redirect)

	log.Println("Server listening on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
