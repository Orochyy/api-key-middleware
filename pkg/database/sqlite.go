package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

const droverName = "sqlite3"

func NewSqliteConnection(database string) *sql.DB {
	return connection(database)
}

func connection(database string) *sql.DB {
	if _, err := os.Stat(database); os.IsNotExist(err) {
		file, err := os.Create(database)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Database was create")

		file.Chmod(os.ModePerm)
		file.Close()
	}

	sqliteDatabase, _ := sql.Open(droverName, database)
	//defer sqliteDatabase.Close() // Defer Closing the database

	return sqliteDatabase
}
