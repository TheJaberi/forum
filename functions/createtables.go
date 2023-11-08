package forum

import (
	"database/sql"
	"fmt"
	"log"
	_"github.com/mattn/go-sqlite3"
)


func CreateTables(title string, post string, username string, password string) {
database, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	usersTable, err2 := database.Prepare("CREATE TABLE IF NOT EXISTS Users (user_id INTEGER PRIMARY KEY, username TEXT, Password TEXT, email TEXT)")
	if err2 != nil {
		log.Fatal(err2)
	}
	postsTable, err3 := database.Prepare("CREATE TABLE IF NOT EXISTS Posts (id INTEGER PRIMARY KEY, Title TEXT, body TEXT, user_id INTEGER NOT NULL)")
	if err3 != nil {
		log.Fatal(err3)
	}
	usersTable.Exec()
	usersTable, _ = database.Prepare("INSERT INTO Users (username, Password) VALUES (?, ?)")
	usersTable.Exec(username, password)
	fmt.Println("test2")
	fmt.Println(title)
	fmt.Println(post)
	postsTable.Exec()
	id, _ := database.Query("SELECT user_id FROM Users")
	var user_id int
	for id.Next(){
	id.Scan(&user_id)}
	postsTable, _ = database.Prepare("INSERT INTO Posts (Title, body, user_id) VALUES (?, ?, ?)")
	postsTable.Exec(title, post, user_id)

}
