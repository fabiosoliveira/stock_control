package product

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func initDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS products (
	 id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	 name TEXT NOT NULL,
	 price FLOAT NOT NULL,
	 stock INTEGER NOT NULL
	);`

	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("Error creating table: %q: %s\n", err, sqlStmt)
	}
}

func init() {
	initDB()
}
