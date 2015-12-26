package main

import (
	"database/sql"
	"net/http"

	"github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
)

// Update do the actual modification of the article in the database.
func ArticleUpdateHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// Ensure the user is logged.
	if sessions.GetSession(r).Get("logged") != true {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var article Article

	err := database.Get(&article, "SELECT * FROM articles WHERE slug = ?", p.ByName("slug"))
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	if err == sql.ErrNoRows {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	// Validate the request's form data for the article.
	var v = NewValidator(r)
	v.NotEmpty("title")
	v.MaxLen("title", 150)
	v.NotEmpty("slug")
	v.MaxLen("slug", 150)
	v.DoesntExists("slug", "articles", "slug", article.Slug)
	v.NotEmpty("tagline")
	v.MaxLen("tagline", 450)
	v.NotEmpty("text")

	// In case of error, forward them to the next request.
	if v.HasErrors() {
		sessions.GetSession(r).AddFlash(v.Errors(), "_errors")
		sessions.GetSession(r).AddFlash(r.Form, "_inputs")
		http.Redirect(w, r, "/article/"+article.Slug+"/edit", http.StatusFound)
		return
	}

	// If everything is fine, update the article in the database.
	_, err = database.Exec("UPDATE articles SET title = ?, slug = ?, tagline = ?, text = ?, tags = ?, updated_at = datetime('now') WHERE id = ?", r.FormValue("title"), r.FormValue("slug"), r.FormValue("tagline"), r.FormValue("text"), r.FormValue("tags"), article.ID)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/article/"+r.FormValue("slug"), http.StatusFound)
}
