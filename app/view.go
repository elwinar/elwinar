package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

const VIEWS_PATH = "resources/views/"

var templates = map[string]*template.Template{}

func view(w http.ResponseWriter, name string, data interface{}) {
	var err error
	
	t, found := templates[name]
	if !found {
		t, err = template.New(name).Funcs(funcs).ParseFiles(filepath.Join(VIEWS_PATH, name + ".html"), filepath.Join(VIEWS_PATH, "app.html"))
		if err != nil {
			panic("unable to parse template index:" + err.Error())
		}
		
		templates[name] = t
	}
	
	t.ExecuteTemplate(w, "app", data)
}
