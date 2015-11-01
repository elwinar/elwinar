package main

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Index displays the home page of the website.
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	timezone, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		timezone = time.FixedZone("Europe/Paris", 1)
	}

	birthDate := time.Date(1990, time.November, 5, 0, 0, 0, 0, timezone)
	now := time.Now()

	age := now.Year() - birthDate.Year()
	if now.YearDay() < birthDate.YearDay(){
		age--
	}

	render(w, r, "index", map[string]interface{}{
		"Title": "Passionate developer",
		"Age":   age,
	})
}
