package dao

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var (
	globalDB *sql.DB
)

func inti() {
	var err error
	globalDB, err = sql.Open("sqlite", "./tmp/week02.db")
	if err != nil {
		panic(err)
	}
}
