package persistsql

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() {
	db, err := sql.Open("sqlite3", "robs.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
