package main

import (
	"net/http"

	"github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
)

// ArticleWriteHandler display the form for writing articles.
func ArticleWriteHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// Ensure the user is logged.
	if sessions.GetSession(r).Get("logged") != true {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	render(w, r, "article_form", map[string]interface{}{
		"Title": "Write",
	})
}
