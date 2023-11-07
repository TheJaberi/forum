package forum

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	_"github.com/mattn/go-sqlite3"
)


func CreateTables(title string, post string) {
	database, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	usersTable, err2 := database.Prepare("CREATE TABLE IF NOT EXISTS Users (id INTEGER PRIMARY KEY, username TEXT, Password TEXT, email TEXT)")
	if err2 != nil {
		log.Fatal(err2)
	}
	postsTable, err3 := database.Prepare("CREATE TABLE IF NOT EXISTS Posts (id INTEGER PRIMARY KEY, Title TEXT, body TEXT)")
	if err3 != nil {
		log.Fatal(err3)
	}
	usersTable.Exec()
	usersTable, _ = database.Prepare("INSERT INTO Users (username, Password) VALUES (?, ?)")
	usersTable.Exec("Nic", "Raboy")
	rows, _ := database.Query("SELECT id, username, Password from Users")
	var id int
	var firstname string
	var lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + firstname + " " + lastname)
	}
	postsTable.Exec()
	postsTable, _ = database.Prepare("INSERT INTO Users (Title, body) VALUES (?, ?)")
	usersTable.Exec(title, post)
	var id2 int
	var head string
	var body string
	rows2, _ := database.Query("SELECT id, Title, body from Posts")
	for rows2.Next() {
		rows2.Scan(&id2, &head, &body)
		fmt.Println(strconv.Itoa(id2) + head + " " + body)
	}
}
