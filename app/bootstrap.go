package main

import (
	"github.com/codegangsta/cli"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

func Bootstrap(context *cli.Context) error {
	var err error
	
	db, err = sqlx.Connect("sqlite3", "storage/database.sqlite")
	if err != nil {
		return err
	}
	
	BASE_URL = context.String("base")
	println("starting with base url", BASE_URL)
	
	return nil
}
