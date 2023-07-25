package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var STRINGPQ = "user=pqgotest dbname=pqgotest sslmode=verify-full"
var DB *sql.DB

func DBconfig() (*sql.DB, error) {
	var err error
	DB, err = sql.Open("postgres", STRINGPQ)

	if err != nil {
		log.Fatal(err)
	}

	return DB, err
}
