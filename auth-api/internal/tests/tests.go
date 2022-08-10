package tests

import (
	"database/sql"
	_ "embed"
	"strings"
)

var (
	//go:embed scripts/clear-script.sql
	clearScript string
)

func FailOnError(err error) {
	if err != nil {
		panic(err)
	}
}
func ClearDB(db *sql.DB) {
	for _, line := range strings.Split(clearScript, "\n") {
		_, execError := db.Exec(line)
		if execError != nil {
			panic(execError)
		}
	}
}
