package main

import (
	"database/sql"
	"net/http"
	"time"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
	"github.com/russross/blackfriday"
)

// Article is the type that holds blog posts.
type Article struct {
	ID          int64     `db:"id"`
	Title       string    `db:"title"`
	Slug        string    `db:"slug"`
	Tagline     string    `db:"tagline"`
	Text        string    `db:"text"`
	Tags        string    `db:"tags"`
	IsPublished bool      `db:"is_published"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	PublishedAt time.Time `db:"published_at"`
}

// MarshalHTML returns the text of the article in HTML format.
func (a Article) MarshalHTML() string {
	return string(blackfriday.MarkdownCommon([]byte(a.Text)))
}

// View handle the displaying of articles.
func View(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

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

// Delete remove articles.
func Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

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

// Edit show the edition form for articles.
func Edit(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

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

	render(w, r, "edit", map[string]interface{}{
		"Title":   "Edit " + article.Title,
		"Article": article,
	})
}

// Update do the actual modification of the article in the database.
func Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

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
	return
}

// Publish set the published flag of an article to true.
func Publish(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

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

// List show all articles available to the user.
func List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var articles []*Article

	// Get the login status of the user.
	logged := sessions.GetSession(r).Get("logged").(bool)

	// If the user isn't logged, only show published articles and order them by publication date.
	// If the user is logged, show all articles and order them by creation date.
	err := database.Select(&articles, "SELECT * FROM articles WHERE is_published = ? ORDER BY published_at DESC", logged)
	if err != nil {
		panic(err)
	}

	render(w, r, "read", map[string]interface{}{
		"Title":    "Read",
		"Articles": articles,
		"Logged":   logged,
	})
}

// Unpublish switch off visibility of an article.
func Unpublish(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

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

// Write display the form for writing articles.
func Write(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// Ensure the user is logged.
	if sessions.GetSession(r).Get("logged") != true {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	render(w, r, "write", map[string]interface{}{
		"Title": "Write",
	})
}

// Create validate and add new articles in the database.
func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

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
