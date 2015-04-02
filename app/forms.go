package main

import (
	"github.com/mholt/binding"
)

type WriteForm struct {
	Title   string
	Slug    string
	Tagline string
	Text    string
	Tags    string
}

func (f *WriteForm) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&f.Title:   binding.Field{
			Form: "title",
			Required: true,
		},
		&f.Slug:    binding.Field{
			Form: "slug",
			Required: true,
		},
		&f.Tagline: binding.Field{
			Form: "tagline",
			Required: true,
		},
		&f.Text:    binding.Field{
			Form: "text",
			Required: true,
		},
		&f.Tags:    "tags",
	}
}
