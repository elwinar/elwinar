package main

import (
	"github.com/codegangsta/cli"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB
var PASSWORD string
var VIEWS string

func Bootstrap(context *cli.Context) error {
	var err error

	db, err = sqlx.Connect("sqlite3", "storage/database.sqlite")
	if err != nil {
		return err
	}

	PASSWORD = context.String("password")
	VIEWS = context.String("views")

	return nil
}
