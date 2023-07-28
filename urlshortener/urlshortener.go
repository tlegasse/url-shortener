package urlshortener

import (
	"database/sql"
	"fmt"
	"net/http"
)

type Shortener struct {
	db *sql.DB
}

func New(db *sql.DB) *Shortener {
	return &Shortener{db: db}
}

func (s *Shortener) Redirect(w http.ResponseWriter, r *http.Request) {
	fmt.Println("shortening")
	http.Redirect(w, r, "https://www.google.com", http.StatusSeeOther)
}

func (s *Shortener) Shorten(w http.ResponseWriter, r *http.Request) {
	fmt.Println("shortening")
}
