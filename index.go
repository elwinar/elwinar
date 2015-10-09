package main

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Index displays the home page of the website.
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	render(w, r, "index", map[string]interface{}{
		"Title": "Passionate developer",
		"Age":   time.Now.Year() - 1990, // I'm born 1990-11-05, and don't care about month precision for this….
	})
}
