package tests

import (
	"database/sql"
	_ "embed"
)

var (
	//go:embed scripts/clear-script.sql
	clearScript string
)

func ClearDB(db *sql.DB) {
	_, execError := db.Exec(clearScript)
	if execError != nil {
		panic(execError)
	}
}
