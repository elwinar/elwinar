package main

import "net/http"

func render(w http.ResponseWriter, name string, data interface{}) {
	t, found := templates[name]
	if !found {
		panic("unknown template " + name)
	}

	t.ExecuteTemplate(w, "app", data)
}
