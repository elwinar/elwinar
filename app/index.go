package main

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	view(w, "index", nil)
}
