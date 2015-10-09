package main

import (
	"crypto/subtle"
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

func LoginHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	render(w, r, "login", map[string]interface{}{
		"Title": "Login",
	})
}

func LoginFormHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if subtle.ConstantTimeEq(int32(len(r.FormValue("password"))), int32(len(configuration.Password))) == 0 {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	if subtle.ConstantTimeCompare([]byte(r.FormValue("password")), []byte(configuration.Password)) == 0 {
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
