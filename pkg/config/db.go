package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB
var (
	POSTGRES_HOST     string = "localhost"
	POSTGRES_PORT     int    = 5432
	POSTGRES_USER     string = "postgres"
	POSTGRES_PASSWORD string = "secretpassword"
	POSTGRES_DBNAME   string = "psql-react"
)

func DBconfig() (*sql.DB, error) {
	stringconn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		POSTGRES_HOST,
		POSTGRES_PORT,
		POSTGRES_USER,
		POSTGRES_PASSWORD,
		POSTGRES_DBNAME,
	)

	var err error
	DB, err = sql.Open("postgres", stringconn)

	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")

	return DB, err
}
