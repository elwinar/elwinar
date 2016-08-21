package main

import (
	"html/template"
	"time"

	"github.com/russross/blackfriday"
)

// Article is the type that holds blog posts.
type Article struct {
	ID          int64     `db:"id"`
	Title       string    `db:"title"`
	Slug        string    `db:"slug"`
	Tagline     string    `db:"tagline"`
	Text        string    `db:"text"`
	Tags        string    `db:"tags"`
	IsPublished bool      `db:"is_published"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	PublishedAt time.Time `db:"published_at"`
}

// MarshalHTML returns the text of the article in HTML format.
func (a Article) MarshalHTML() template.HTML {
	return template.HTML(blackfriday.MarkdownCommon([]byte(a.Text)))
}
