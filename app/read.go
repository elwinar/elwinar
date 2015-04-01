package main

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
)

func ReadHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var articles []Article
	
	err := db.Select(&articles, "SELECT * FROM articles ORDER BY created_at DESC")
	if err != nil {
		panic(err)
	}
	
	view(w, "read", articles)
}
