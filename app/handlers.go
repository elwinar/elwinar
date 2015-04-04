package main

import (
	"crypto/subtle"
	"fmt"
	"github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
	"github.com/sourcegraph/sitemap"
	"net/http"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	article := FindArticle(p.ByName("slug"))
	if article == nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	if sessions.GetSession(r).Get("logged") != true && !article.IsPublished {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	render(w, r, "article", article)
}

func EditArticleHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	article := FindArticle(p.ByName("slug"))
	if article == nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	render(w, r, "edit", article)
}

func EditArticleFormHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	article := FindArticle(p.ByName("slug"))
	if article == nil {
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

	_, err := db.Exec("UPDATE articles SET title = ?, slug = ?, tagline = ?, text = ?, tags = ?, updated_at = datetime('now') WHERE id = ?", r.FormValue("title"), r.FormValue("slug"), r.FormValue("tagline"), r.FormValue("text"), r.FormValue("tags"), article.ID)
	if err != nil {
		panic(err)
	}
	
	http.Redirect(w, r, "/article/"+r.FormValue("slug"), http.StatusFound)
	return
}

func DeleteArticleHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_, err := db.Exec("DELETE FROM articles WHERE slug = ?", p.ByName("slug"))
	if err != nil {
		panic(err)
	}
	
	http.Redirect(w, r, "/read", http.StatusFound)
}

func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	render(w, r, "index", nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	render(w, r, "login", nil)
}

func LoginFormHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if subtle.ConstantTimeEq(int32(len(r.FormValue("password"))), int32(len(PASSWORD))) == 0 {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	if subtle.ConstantTimeCompare([]byte(r.FormValue("password")), []byte(PASSWORD)) == 0 {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	sessions.GetSession(r).Set("logged", true)
	http.Redirect(w, r, "/read", http.StatusFound)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sessions.GetSession(r).Clear()
	http.Redirect(w, r, r.Referer(), http.StatusFound)
}

func PublishArticleHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_, err := db.Exec("UPDATE articles SET is_published = ?, published_at = datetime('now') WHERE slug = ?", true, p.ByName("slug"))
	if err != nil {
		panic(err)
	}
	
	http.Redirect(w, r, r.Referer(), http.StatusFound)
}

func ReadHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	articles := AllArticles()
	render(w, r, "read", articles)
}

func SitemapHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	articles := PublishedArticles()
	
	var urlset sitemap.URLSet
	urlset.URLs = []sitemap.URL{
		{
			Loc: fmt.Sprintf("%s/", BASE),
			ChangeFreq: sitemap.Yearly,
		},
		{
			Loc: fmt.Sprintf("%s/read", BASE),
			ChangeFreq: sitemap.Weekly,
		},
	}
	
	for _, a := range articles {
		urlset.URLs = append(urlset.URLs, sitemap.URL{
			Loc: fmt.Sprintf("%s/article/%s", BASE, a.Slug),
			LastMod: &a.UpdatedAt,
			ChangeFreq: sitemap.Monthly,
		})
	}
	
	raw, err := sitemap.Marshal(&urlset)
	if err != nil {
		panic(err)
	}
	
	w.Write(raw)
}

func UnpublishArticleHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_, err := db.Exec("UPDATE articles SET is_published = ? WHERE slug = ?", false, p.ByName("slug"))
	if err != nil {
		panic(err)
	}
	
	http.Redirect(w, r, r.Referer(), http.StatusFound)
}

func WriteHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	render(w, r, "write", nil)
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

	_, err := db.Exec("INSERT INTO articles (title, slug, tagline, text, tags, created_at, updated_at, is_published) VALUES (?, ?, ?, ?, ?, datetime('now'), datetime('now'), ?)", r.FormValue("title"), r.FormValue("slug"), r.FormValue("tagline"), r.FormValue("text"), r.FormValue("tags"), false)
	if err != nil {
		panic(err)
	}
	
	http.Redirect(w, r, "/article/"+r.FormValue("slug"), http.StatusFound)
	return
}
