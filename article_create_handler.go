package main

import (
	"net/http"

	"github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
)

// Create validate and add new articles in the database.
func ArticleCreateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// Ensure the user is logged.
	if sessions.GetSession(r).Get("logged") != true {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Validate the request form data.
	var v = NewValidator(r)
	v.NotEmpty("title")
	v.MaxLen("title", 150)
	v.NotEmpty("slug")
	v.MaxLen("slug", 150)
	v.DoesntExists("slug", "articles", "slug")
	v.NotEmpty("tagline")
	v.MaxLen("tagline", 450)
	v.NotEmpty("text")

	// In case of error, return the user to the previous page with a listing of
	// the errors and inputs.
	if v.HasErrors() {
		sessions.GetSession(r).AddFlash(v.Errors(), "_errors")
		sessions.GetSession(r).AddFlash(r.Form, "_inputs")
		http.Redirect(w, r, "/write", http.StatusFound)
		return
	}

	// Insert the new article in the database.
	_, err := database.Exec("INSERT INTO articles (title, slug, tagline, text, tags, is_published, created_at, updated_at, published_at) VALUES (?, ?, ?, ?, ?, ?, datetime('now'), datetime('now'), datetime('now'))", r.FormValue("title"), r.FormValue("slug"), r.FormValue("tagline"), r.FormValue("text"), r.FormValue("tags"), false)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/article/"+r.FormValue("slug"), http.StatusFound)
}
