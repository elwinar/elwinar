package main

import (
	"database/sql"
	"net/http"
	"time"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
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

func ArticleHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var article Article

	err := database.Get(&article, "SELECT * FROM articles WHERE slug = ?", p.ByName("slug"))
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	if err == sql.ErrNoRows {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	if sessions.GetSession(r).Get("logged") != true && !article.IsPublished {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	render(w, r, "article", map[string]interface{}{
		"Title":   article.Title,
		"Article": article,
	})
}

func DeleteArticleHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_, err := database.Exec("DELETE FROM articles WHERE slug = ?", p.ByName("slug"))
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/read", http.StatusFound)
}

func EditArticleHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var article Article

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

func EditArticleFormHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var article Article

	err := database.Get(&article, "SELECT * FROM articles WHERE slug = ?", p.ByName("slug"))
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	if err == sql.ErrNoRows {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	var v = NewValidator(r)
	v.NotEmpty("title")
	v.MaxLen("title", 150)
	v.NotEmpty("slug")
	v.MaxLen("slug", 150)
	v.DoesntExists("slug", "articles", "slug", article.Slug)
	v.NotEmpty("tagline")
	v.MaxLen("tagline", 450)
	v.NotEmpty("text")

	if v.HasErrors() {
		sessions.GetSession(r).AddFlash(v.Errors(), "_errors")
		sessions.GetSession(r).AddFlash(r.Form, "_inputs")
		http.Redirect(w, r, "/article/"+article.Slug+"/edit", http.StatusFound)
		return
	}

	_, err = database.Exec("UPDATE articles SET title = ?, slug = ?, tagline = ?, text = ?, tags = ?, updated_at = datetime('now') WHERE id = ?", r.FormValue("title"), r.FormValue("slug"), r.FormValue("tagline"), r.FormValue("text"), r.FormValue("tags"), article.ID)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/article/"+r.FormValue("slug"), http.StatusFound)
	return
}

func PublishArticleHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_, err := database.Exec("UPDATE articles SET is_published = ?, published_at = datetime('now') WHERE slug = ?", true, p.ByName("slug"))
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, r.Referer(), http.StatusFound)
}

func ReadHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var articles []*Article

	err := database.Select(&articles, "SELECT * FROM articles ORDER BY published_at DESC")
	if err != nil {
		panic(err)
	}

	render(w, r, "read", map[string]interface{}{
		"Title":    "Read",
		"Articles": articles,
	})
}

func UnpublishArticleHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_, err := database.Exec("UPDATE articles SET is_published = ? WHERE slug = ?", false, p.ByName("slug"))
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, r.Referer(), http.StatusFound)
}

func WriteHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	render(w, r, "write", map[string]interface{}{
		"Title": "Write",
	})
}

func WriteFormHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var v = NewValidator(r)
	v.NotEmpty("title")
	v.MaxLen("title", 150)
	v.NotEmpty("slug")
	v.MaxLen("slug", 150)
	v.DoesntExists("slug", "articles", "slug")
	v.NotEmpty("tagline")
	v.MaxLen("tagline", 450)
	v.NotEmpty("text")

	if v.HasErrors() {
		sessions.GetSession(r).AddFlash(v.Errors(), "_errors")
		sessions.GetSession(r).AddFlash(r.Form, "_inputs")
		http.Redirect(w, r, "/write", http.StatusFound)
		return
	}

	_, err := database.Exec("INSERT INTO articles (title, slug, tagline, text, tags, is_published, created_at, updated_at, published_at) VALUES (?, ?, ?, ?, ?, ?, datetime('now'), datetime('now'), datetime('now'))", r.FormValue("title"), r.FormValue("slug"), r.FormValue("tagline"), r.FormValue("text"), r.FormValue("tags"), false)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/article/"+r.FormValue("slug"), http.StatusFound)
	return
}
