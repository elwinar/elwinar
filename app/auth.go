package main

import (
	"net/http"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
)

func Auth(handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if sessions.GetSession(r).Get("logged") != true {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		handle(w, r, p)
	}
}
