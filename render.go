package main

import (
	"net/http"
	"net/url"

	sessions "github.com/goincremental/negroni-sessions"
)

func render(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	var rawErrs = sessions.GetSession(r).Flashes("_errors")
	var errs []string
	if len(rawErrs) != 0 {
		errs = rawErrs[0].([]string)
	}

	var rawInputs = sessions.GetSession(r).Flashes("_inputs")
	var inputs = new(url.Values)
	if len(rawInputs) != 0 {
		inputs = rawInputs[0].(*url.Values)
	}

	t, found := templates[name]
	if !found {
		panic("unknown template " + name)
	}

	t.ExecuteTemplate(w, "app", data)
}
