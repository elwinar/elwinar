package main

import (
	"database/sql"
	"net/http"

	"github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
)

// Edit show the edition form for articles.
func ArticleEditHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// Ensure the user is logged
	if sessions.GetSession(r).Get("logged") != true {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var article Article

	// Get the article from the database.
	err := database.Get(&article, "SELECT * FROM articles WHERE slug = ?", p.ByName("slug"))
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	if err == sql.ErrNoRows {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	render(w, r, "article_edit", map[string]interface{}{
		"Title":   "Edit " + article.Title,
		"Article": article,
	})
}
