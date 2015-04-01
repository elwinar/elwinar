package main

import (
	"github.com/russross/blackfriday"
	"html/template"
	"strconv"
	"time"
)

var funcs = template.FuncMap{
	"dateFormat": DateFormatHelper,
	"markdown": MarkdownHelper,
	"age": AgeHelper,
	"baseURL": BaseUrlHelper,
}

func DateFormatHelper(format string, date time.Time) string {
	return date.Format(format)
}

func MarkdownHelper(text string) template.HTML {
	return template.HTML(blackfriday.MarkdownCommon([]byte(text)))
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

var BASE_URL string
func BaseUrlHelper() string {
	return BASE_URL
}
