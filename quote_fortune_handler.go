package main

import (
	"database/sql"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Roll a new fortune.
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

	render(w, r, "fortune", map[string]interface{}{
		"Title": "Fortune",
		"Quote": quote,
	})
}
