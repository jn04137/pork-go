package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var DB *sql.DB

//var connStr string = "postgres://knostash_user:example@localhost:5432/knostash_db?sslmode=disable"
var connStr string = os.Getenv("DB_CONN")

func Open() error {
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return err
}

func Close() error {
	return DB.Close()
}
