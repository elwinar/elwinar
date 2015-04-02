package main

import (
	"time"
)

type Article struct {
	ID int64 `db:"id"`
	Title string `db:"title"`
	Tagline string `db:"tagline"`
	Slug string `db:"slug"`
	Text string `db:"text"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Published bool `db:"published"`
}
