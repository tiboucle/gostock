package controller

import (
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

var DB *sql.DB

func WriteDefaultResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Unsupported Request"))
}

func WriteJSON(w http.ResponseWriter, i interface{}) {
	result, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(result)
}
