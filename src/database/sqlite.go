package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const (
	apiURL     = "https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=IBM&apikey=your_api_key"
	dbFileName = "./addons/apollo.db"
)

func openConnection() *sql.DB {
	db, err := sql.Open("sqlite3", dbFileName)

	if err != nil {
		panic(err)
	}

	return db
}

func storeData(db *sql.DB) error {
	db = openConnection()

	return fmt.Errorf("")
}
