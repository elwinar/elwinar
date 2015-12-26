package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

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
