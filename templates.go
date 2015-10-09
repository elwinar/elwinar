package main

import (
	"html/template"
	"net/http"
)

// templates contains all the templates parsed at run-time.
var templates map[string]*template.Template

// init parse the layout template and the view templates.
func init() {
	templates = make(map[string]*template.Template)

	raw, err := Asset("views/layout.html")
	if err != nil {
		panic(err)
	}

	layout, err := template.New("layout").Parse(string(raw))
	if err != nil {
		panic(err)
	}

	for _, template := range []string{
		"article_view",
		"article_form",
		"article_list",
		"article_edit",
		"quote",
		"index",
		"login",
	} {
		t, _ := layout.Clone()

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
func render(w http.ResponseWriter, name string, data interface{}) {
	t, found := templates[name]
	if !found {
		panic("unknown template " + name)
	}

	t.ExecuteTemplate(w, "app", data)
}
