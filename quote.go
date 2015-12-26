package main

type Quote struct {
	ID     int64  `db:"id"`
	Text   string `db:"text"`
	Author string `db:"author"`
}
