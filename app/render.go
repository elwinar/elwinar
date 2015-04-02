package main

import (
	"fmt"
	"github.com/goincremental/negroni-sessions"
	"github.com/russross/blackfriday"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func render(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	var session = sessions.GetSession(r)
	
	t, err := template.New(name).Funcs(template.FuncMap{
		"age": func(birthdate string) (string, error) {
			var birth, err = time.Parse("2006-01-02", birthdate)
			if err != nil {
				return "", err
			}
			
			var now = time.Now().UTC()
			if now.YearDay() < birth.YearDay() {
				return strconv.Itoa(now.Year() - birth.Year() - 1), nil
			}
			
			return strconv.Itoa(now.Year() - birth.Year()), nil
		},
		"base": func() string {
			return fmt.Sprintf("http://%s", r.Host)
		},
		"dateFormat": func(format string, date time.Time) string {
			return date.Format(format)
		},
		"errors": func() []interface{} {
			return session.Flashes("_errors")
		},
		"env": func(name string) string {
			return os.Getenv(name)
		},
		"logged": func() interface{} {
			return session.Get("logged")
		},
		"markdown": func(text string) template.HTML {
			return template.HTML(blackfriday.MarkdownCommon([]byte(text)))
		},
		"session": func(name string) interface{} {
			return session.Get(name)
		},
	}).ParseFiles(filepath.Join(VIEWS, "app.html"), filepath.Join(VIEWS, name + ".html"))
	if err != nil {
		panic("unable to parse template "+name+":" + err.Error())
	}
	
	t.ExecuteTemplate(w, "app", data)
}
