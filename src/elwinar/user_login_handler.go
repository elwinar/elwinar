package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Login display the login form.
func UserLoginHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	render(w, r, "login", map[string]interface{}{
		"Title": "Login",
	})
}
