package main

import (
	"net/http"
	"time"

	"github.com/elwinar/heirloom"
	sessions "github.com/goincremental/negroni-sessions"
)

// The templates are stored into another directory.
//go:generate go-bindata -nomemcopy -prefix "../" -pkg $GOPACKAGE -o views.go ../views

// templates holds the collection of templates for the whole website.
var templates *heirloom.Heirloom

// init parse all the embeded templates and add them to the collection.
func init() {
	templates = heirloom.New()
	templates.Funcs(heirloom.FuncMap{
		"Format": func(format string, date time.Time) string {
			return date.Format(format)
		},
	})

	for _, template := range []string{
		"layout",
		"article",
		"article_list",
		"article_form",
		"article_edit",
		"index",
		"fortune",
		"login",
	} {
		raw, err := Asset("views/" + template + ".html")
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
	data["Configuration"] = configuration

	out, err := templates.Render(name, data)
	if err != nil {
		http.Error(w, "error while generating the page:"+err.Error(), http.StatusInternalServerError)
	}

	w.Write([]byte(out))
}
