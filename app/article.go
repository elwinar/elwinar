package main

import (
	"database/sql"
	"time"
)

type Article struct {
	ID          int64     `db:"id"`
	Title       string    `db:"title"`
	Slug        string    `db:"slug"`
	Tagline     string    `db:"tagline"`
	Text        string    `db:"text"`
	Tags        string    `db:"tags"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	IsPublished bool      `db:"is_published"`
}

func FindArticle(slug string) (*Article) {
	var article Article
	err := db.Get(&article, "SELECT * FROM articles WHERE slug = ?", slug)
	
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	
	if err == sql.ErrNoRows {
		return nil
	}
	
	return &article
}
