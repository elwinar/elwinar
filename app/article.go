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
	IsPublished bool      `db:"is_published"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	PublishedAt time.Time `db:"published_at"`
}

func FindArticle(slug string) (*Article) {
	var article Article
	err := db.Get(&article, "SELECT id, title, slug, tagline, text, tags, is_published, created_at, updated_at, published_at FROM articles WHERE slug = ?", slug)
	
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	
	if err == sql.ErrNoRows {
		return nil
	}
	
	return &article
}

func AllArticles() ([]*Article) {
	var articles []*Article

	err := db.Select(&articles, "SELECT id, title, slug, tagline, text, tags, is_published, created_at, updated_at, published_at FROM articles ORDER BY published_at DESC")
	if err != nil {
		panic(err)
	}

	return articles
}

func PublishedArticles() ([]*Article) {
	var articles []*Article
	
	err := db.Select(&articles, "SELECT id, title, slug, tagline, text, tags, is_published, created_at, updated_at, published_at FROM articles WHERE is_published = ? ORDER BY updated_at DESC", true)
	if err != nil {
		panic(err)
	}
	
	return articles
}
