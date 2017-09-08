package model

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func SetDatabase(database *sql.DB) {
	db = database
}
