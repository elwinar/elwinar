package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// Configuration holds the configuration for the website: which port to listen on,
// which database to use, etc.
type Configuration struct {
	Database       string
	Debug          bool
	Password       string
	Port           int
	Public         string
	Secret         string
	AnalyticsToken string
}

func ConfigurationHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	raw, _ := json.Marshal(configuration)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(raw))
}

func EnvironmentHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	response := strings.Join(os.Environ(), "\n")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
