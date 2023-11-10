package forum

import (
	"database/sql"
	"fmt"
	"log"
	_"github.com/mattn/go-sqlite3"
)


func CreateTables() {
fmt.Println(66)
	Database, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	// _, errpragma := Database.Exec("PRAGMA foreign_keys = ON")
	// if errpragma != nil {
	// 	log.Fatal(errpragma)
	// }
	usertable, err2 := Database.Prepare("CREATE TABLE IF NOT EXISTS Users (id INTEGER PRIMARY KEY, username TEXT NOT NULL, Password TEXT NOT NULL, email TEXT)")
	if err2 != nil {
		log.Fatal(err2)
	}
	usertable.Exec()
	postsTable, err4 := Database.Prepare("CREATE TABLE IF NOT EXISTS Posts (id INTEGER PRIMARY KEY, Title TEXT, body TEXT)")
	if err4 != nil {
		log.Fatal(err4)
	}
	postsTable.Exec()
	defer Database.Close()
}