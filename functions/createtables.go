package forum

import (
	"database/sql"
	// "fmt"
	"log"
	_"github.com/mattn/go-sqlite3"
)


func CreateTables(title string, post string, username string, password string) {
database, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	_, errpragma := database.Exec("PRAGMA foreign_keys = ON")
	if errpragma != nil {
		log.Fatal(errpragma)
	}
	usersTable, err2 := database.Prepare("CREATE TABLE IF NOT EXISTS Users (user_id INTEGER PRIMARY KEY, username TEXT NOT NULL, Password TEXT NOT NULL, email TEXT)")
	if err2 != nil {
		log.Fatal(err2)
	}
	postsTable, err3 := database.Prepare("CREATE TABLE IF NOT EXISTS Posts (id INTEGER PRIMARY KEY, Title TEXT, body TEXT, user_id INTEGER, user_id REFERENCES Users(user_id))")
	if err3 != nil {
		log.Fatal(err3)
	}
	if username != "" && password != ""{
	usersTable.Exec()
	usersTable, _ = database.Prepare("INSERT INTO Users (username, Password) VALUES (?, ?)")
	usersTable.Exec(username, password)}
	// fmt.Println("test2")
	// fmt.Println(title)
	// fmt.Println(post)
	postsTable.Exec()
	// id, _ := database.Query("SELECT user_id, username FROM Users")
// 	var user_id int
// 	var useridfinal int
// 	var username2 string
// 	for id.Next(){
// 	id.Scan(&user_id, &username2)
// 	fmt.Println(username2)
// 	if username2 == "user"{
// useridfinal = user_id}}
// row := database.QueryRow(`INSERT INTO Posts (user_id) VALUES () RETURNING id`)
	postsTable, _ = database.Prepare("INSERT INTO Posts (Title, body, user_id) VALUES (?, ?, ?)")
	postsTable.Exec(title, post)
}