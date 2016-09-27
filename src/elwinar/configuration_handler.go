package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ConfigurationHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	raw, _ := json.Marshal(configuration)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(raw))
}
