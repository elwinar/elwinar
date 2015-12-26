package main

import (
	"net/http"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
)

// Logout de-authenticate the user.
func UserLogoutHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	sessions.GetSession(r).Clear()
	http.Redirect(w, r, r.Referer(), http.StatusFound)
}
