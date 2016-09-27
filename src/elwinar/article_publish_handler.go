package main

import (
	"net/http"

	"github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
)

// ArticlePublishHandler set the published flag of an article to true.
func ArticlePublishHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// Ensure the user is logged.
	if sessions.GetSession(r).Get("logged") != true {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Switch the flag on.
	_, err := database.Exec("UPDATE articles SET is_published = ?, published_at = datetime('now') WHERE slug = ?", true, p.ByName("slug"))
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, r.Referer(), http.StatusFound)
}
