package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	render(w, r, "index", map[string]interface{}{
		"Title": "Passionate developer",
	})
}
