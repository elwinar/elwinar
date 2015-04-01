package main

import (
	"database/sql"
    "github.com/julienschmidt/httprouter"
    "net/http"
	"time"
)

type Article struct {
	ID int64 `db:"id"`
	Title string `db:"title"`
	TagLine string `db:"tagline"`
	Slug string `db:"slug"`
	Text string `db:"text"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Published bool `db:"published"`
}

func ArticleHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var article Article
	
	err := db.Get(&article, "SELECT * FROM articles WHERE slug = ?", p.ByName("slug"))
	if err == sql.ErrNoRows {
		http.Error(w, "Page not found", 404)
		return
	}
	
	if err != nil {
		panic(err)
	}
	
	if !article.Published {
		http.Error(w, "Page not found", 404)
		return
	}
	
	view(w, "article", article)
}
