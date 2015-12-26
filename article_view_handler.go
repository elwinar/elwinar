package main

import (
	"database/sql"
	"net/http"

	"github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
)

// View handle the displaying of articles.
func ArticleViewHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var article Article

	err := database.Get(&article, "SELECT * FROM articles WHERE slug = ?", p.ByName("slug"))
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	if err == sql.ErrNoRows {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	// Reject the request if the user isn't logged and that the article isn't published.
	if sessions.GetSession(r).Get("logged") != true && !article.IsPublished {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	render(w, r, "article", map[string]interface{}{
		"Title":   article.Title,
		"Article": article,
	})
}
