package main

import (
	"net/http"

	"github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
)

// Delete remove articles.
func ArticleDeleteHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// Ensure the user is logged.
	if sessions.GetSession(r).Get("logged") != true {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Remote the article from the database.
	_, err := database.Exec("DELETE FROM articles WHERE slug = ?", p.ByName("slug"))
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/read", http.StatusFound)
}
