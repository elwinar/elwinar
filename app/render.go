package main

import (
	"encoding/gob"
	"github.com/goincremental/negroni-sessions"
	"github.com/russross/blackfriday"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func init() {
	gob.Register(&url.Values{})
}

func AgeHelper(birthdate string) (string, error) {
	var birth, err = time.Parse("2006-01-02", birthdate)
	if err != nil {
		return "", err
	}

	var now = time.Now().UTC()
	if now.YearDay() < birth.YearDay() {
		return strconv.Itoa(now.Year() - birth.Year() - 1), nil
	}

	return strconv.Itoa(now.Year() - birth.Year()), nil
}

func BaseHelper() string {
	return BASE
}

func DateFormatHelper(format string, date time.Time) string {
	return date.Format(format)
}

func EnvHelper(name string) string {
	return os.Getenv(name)
}

func MarkdownHelper(text string) template.HTML {
	return template.HTML(blackfriday.MarkdownCommon([]byte(text)))
}

func GoogleAnalyticsIDHelper() string {
	return GOOGLE_ANALYTICS_ID
}

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

	t, err := template.New(name).Funcs(template.FuncMap{
		"Age":        AgeHelper,
		"Base":       BaseHelper,
		"DateFormat": DateFormatHelper,
		"Errors": func() []string {
			return errs
		},
		"Env":               EnvHelper,
		"GoogleAnalyticsID": GoogleAnalyticsIDHelper,
		"Input": func() *url.Values {
			return inputs
		},
		"Logged": func() interface{} {
			return sessions.GetSession(r).Get("logged")
		},
		"Markdown": MarkdownHelper,
		"Session": func(name string) interface{} {
			return sessions.GetSession(r).Get(name)
		},
	}).ParseFiles(filepath.Join(VIEWS, "app.html"), filepath.Join(VIEWS, name+".html"))
	if err != nil {
		panic("unable to parse template " + name + ":" + err.Error())
	}

	t.ExecuteTemplate(w, "app", data)
}
