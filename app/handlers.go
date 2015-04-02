package main

import (
	"crypto/subtle"
	"database/sql"
	"fmt"
	"github.com/goincremental/negroni-sessions"
    "github.com/julienschmidt/httprouter"
	"github.com/mholt/binding"
    "net/http"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var article Article
	
	err := db.Get(&article, "SELECT * FROM articles WHERE slug = ?", p.ByName("slug"))
	if err == sql.ErrNoRows {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
	
	if err != nil {
		panic(err)
	}
	
	if !article.Published {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
	
	render(w, r, "article", article)
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
	http.Redirect(w, r, "/", http.StatusFound)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sessions.GetSession(r).Clear()
	http.Redirect(w, r, "/", http.StatusFound)
}

func ReadHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var articles []Article
	
	err := db.Select(&articles, "SELECT * FROM articles ORDER BY created_at DESC")
	if err != nil {
		panic(err)
	}
	
	render(w, r, "read", articles)
}

func WriteHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	render(w, r, "write", nil)
}

func WriteFormHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var form WriteForm
	
	errs := binding.Bind(r, &form)
	if len(errs) != 0 {
		http.Redirect(w, r, "/write", http.StatusFound)
		return
	}
	
	fmt.Fprintf(w, "%v", form)
}
