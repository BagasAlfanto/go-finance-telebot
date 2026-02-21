package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "./database/finance.db")
	if err != nil {
		log.Panic(err)
	}

	createTables := `
	CREATE TABLE IF NOT EXISTS pemasukan (
		id         INTEGER PRIMARY KEY AUTOINCREMENT,
		nominal    INTEGER NOT NULL,
		deskripsi  TEXT    NOT NULL,
		kategori   TEXT    NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS pengeluaran (
		id         INTEGER PRIMARY KEY AUTOINCREMENT,
		nominal    INTEGER NOT NULL,
		deskripsi  TEXT    NOT NULL,
		kategori   TEXT    NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = DB.Exec(createTables)
	if err != nil {
		log.Panic(err)
	}
}
