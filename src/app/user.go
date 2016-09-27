package main

import (
	"crypto/subtle"
	"net/http"

	"github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
)

// Authenticate check the user credentials.
func UserAuthenticateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

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

// Login display the login form.
func UserLoginHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	render(w, r, "login", map[string]interface{}{
		"Title": "Login",
	})
}

// Logout de-authenticate the user.
func UserLogoutHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	sessions.GetSession(r).Clear()
	http.Redirect(w, r, r.Referer(), http.StatusFound)
}
