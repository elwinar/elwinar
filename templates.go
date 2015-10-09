package main

import (
	"html/template"
)

var templates map[string]*template.Template

func init() {
	templates = make(map[string]*template.Template)

	raw, err := Asset("views/app.html")
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

		raw, err := Asset("../views/" + template + ".html")
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
