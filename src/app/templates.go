package main

import (
	"encoding/gob"
	"net/http"
	"net/url"
	"time"

	"github.com/elwinar/heirloom"
	sessions "github.com/goincremental/negroni-sessions"
)

// templates holds the collection of templates for the whole website.
var templates *heirloom.Heirloom

// init parse all the embeded templates and add them to the collection.
func init() {
	gob.Register(&url.Values{})

	templates = heirloom.New()
	templates.Funcs(heirloom.FuncMap{
		"Format": func(format string, date time.Time) string {
			return date.Format(format)
		},
	})

	for _, template := range []string{
		"layout",
		"index",
		"login",
		"article_view",
		"article_list",
		"article_form",
		"article_edit",
		"quote_view",
		"quote_list",
	} {
		raw, err := Asset("src/views/" + template + ".html")
		if err != nil {
			panic(err)
		}

		err = templates.Parse(template, string(raw))
		if err != nil {
			panic(err)
		}
	}
}

// render write the given template on a response writer.
func render(w http.ResponseWriter, r *http.Request, name string, data map[string]interface{}) {
	data["Logged"] = sessions.GetSession(r).Get("logged")
	data["Request"] = r
	data["Configuration"] = configuration

	var rawErrors = sessions.GetSession(r).Flashes("_errors")
	if len(rawErrors) != 0 {
		data["Errors"] = rawErrors[0].([]string)
	}

	var rawInputs = sessions.GetSession(r).Flashes("_inputs")
	if len(rawInputs) != 0 {
		data["Inputs"] = rawInputs[0].(*url.Values)
	}

	out, err := templates.Render(name, data)
	if err != nil {
		http.Error(w, "error while generating the page:"+err.Error(), http.StatusInternalServerError)
	}

	w.Write([]byte(out))
}
