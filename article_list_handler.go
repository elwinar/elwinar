package main

import (
	"net/http"

	"github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
)

// List show all articles available to the user.
func ArticleListHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var articles []*Article

	// If the user isn't logged, only show published articles and order them by publication date.
	// If the user is logged, show all articles and order them by creation date.
	var err error
	if sessions.GetSession(r).Get("logged") != true {
		err = database.Select(&articles, "SELECT * FROM articles WHERE is_published = ? ORDER BY published_at DESC", true)
	} else {
		err = database.Select(&articles, "SELECT * FROM articles ORDER BY published_at DESC")
	}
	if err != nil {
		panic(err)
	}

	render(w, r, "article_list", map[string]interface{}{
		"Title":    "Read",
		"Articles": articles,
	})
}
