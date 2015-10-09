package main

import (
	"html/template"
	"net/http"
	"time"

	sessions "github.com/goincremental/negroni-sessions"
)

// templates contains all the templates parsed at run-time.
var templates map[string]*template.Template

// helpers define new functions available in the templates.
var helpers = template.FuncMap{
	// Format return the string corresponding to the given format for the given date.
	"Format": func(format string, date time.Time) string {
		return date.Format(format)
	},
}

// init parse the layout template and the view templates.
func init() {
	templates = make(map[string]*template.Template)

	raw, err := Asset("views/layout.html")
	if err != nil {
		panic(err)
	}

	layout, err := template.New("layout").Funcs(helpers).Parse(string(raw))
	if err != nil {
		panic(err)
	}

	for _, template := range []string{
		"article",
		"article_list",
		"article_form",
		"article_edit",
		"index",
		"quote",
		"login",
	} {
		t, err := layout.Clone()

		raw, err := Asset("views/" + template + ".html")
		if err != nil {
			panic(err)
		}

		t, err = t.Parse(string(raw))
		if err != nil {
			panic(err)
		}

		templates[template] = t
	}
}

// render write the given template on a response writer.
func render(w http.ResponseWriter, r *http.Request, name string, data map[string]interface{}) {
	t, found := templates[name]
	if !found {
		panic("unknown template " + name)
	}

	data["Logged"] = sessions.GetSession(r).Get("logged")
	data["Configuration"] = configuration

	t.ExecuteTemplate(w, "app", data)
}
