package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type Quote struct {
	ID     int64  `db:"id"`
	Text   string `db:"text"`
	Author string `db:"author"`
}

// QuoteFortuneHandler display a random quote.
func QuoteFortuneHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var quote Quote

	err := database.Get(&quote, "SELECT * FROM quotes WHERE id >= (select ABS(RANDOM()) % MAX(id) + 1 FROM quotes) LIMIT 1")
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	if err == sql.ErrNoRows {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	render(w, r, "quote_view", map[string]interface{}{
		"Title": "Fortune",
		"Quote": quote,
	})
}

// QuoteListHandler display a quote.
func QuoteListHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var quotes []Quote

	err := database.Select(&quotes, "SELECT * FROM quotes ORDER BY id")
	if err != nil {
		panic(err)
	}

	render(w, r, "quote_list", map[string]interface{}{
		"Title":  "Quotes",
		"Quotes": quotes,
	})
}

// QuoteViewHandler display a quote.
func QuoteViewHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var quote Quote

	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = database.Get(&quote, "SELECT * FROM quotes WHERE id = ?", id)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	if err == sql.ErrNoRows {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	render(w, r, "quote_view", map[string]interface{}{
		"Title": "Fortune",
		"Quote": quote,
	})
}
