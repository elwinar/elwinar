package main

import (
	"net/http"

	"github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
)

// Unpublish switch off visibility of an article.
func ArticleUnpublishHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// Ensure the user is logged.
	if sessions.GetSession(r).Get("logged") != true {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Switch the flag off.
	_, err := database.Exec("UPDATE articles SET is_published = ? WHERE slug = ?", false, p.ByName("slug"))
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, r.Referer(), http.StatusFound)
}
