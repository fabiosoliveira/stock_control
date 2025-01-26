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

	// inserir somente uma vez
	// contar quantidade de linhas na tabela
	var count int
	err = DB.QueryRow("SELECT COUNT(*) FROM products").Scan(&count)
	if err != nil {
		log.Fatalf("Error counting rows: %q\n", err)
	}

	if count == 0 {
		_, err = DB.Exec("INSERT INTO products (name, price, stock) VALUES (?, ?, ?)", "A", 10.0, 10)
		if err != nil {
			log.Fatalf("Error inserting data: %q\n", err)
		}

		_, err = DB.Exec("INSERT INTO products (name, price, stock) VALUES (?, ?, ?)", "B", 20.0, 20)
		if err != nil {
			log.Fatalf("Error inserting data: %q\n", err)
		}
	}
}

func init() {
	initDB()
}
