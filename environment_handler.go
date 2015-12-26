package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func EnvironmentHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	response := strings.Join(os.Environ(), "\n")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
