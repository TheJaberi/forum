package forum

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func DataBase() {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var query string
	query = `CREATE TABLE IF NOT EXISTS users (
		user_id INTEGER NOT NULL,
		user_name CHAR(10) NOT NULL UNIQUE,
		user_email CHAR(25) NOT NULL UNIQUE,
		user_pass CHAR(25) NOT NULL,
		user_type TEXT NOT NULL DEFAULT member,
		time_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY("user_id" AUTOINCREMENT)
	);`
	usersTable, _ := db.Prepare(query)
	usersTable.Exec()
	query = `CREATE TABLE IF NOT EXISTS posts (
		post_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		title TEXT NOT NULL,
		img_url TEXT,
		body TEXT NOT NULL,
		time_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY("post_id" AUTOINCREMENT),
		FOREIGN KEY("user_id") REFERENCES users("user_id")
	)`
	postsTable, _ := db.Prepare(query)
	postsTable.Exec()
	query = `CREATE TABLE IF NOT EXISTS interaction_posts (
		post_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		interaction TEXT NOT NULL,
		FOREIGN KEY("post_id") REFERENCES posts("post_id"),
		FOREIGN KEY("user_id") REFERENCES users("user_id")
	)`
	interaction_postsTable, _ := db.Prepare(query)
	interaction_postsTable.Exec()
	query = `CREATE TABLE IF NOT EXISTS comments (
		comment_id INTEGER NOT NULL,
		post_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		body TEXT NOT NULL,
		time_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY("comment_id" AUTOINCREMENT),
		FOREIGN KEY("post_id") REFERENCES posts("post_id"),
		FOREIGN KEY("user_id") REFERENCES users("user_id")
	)`
	commentsTable, _ := db.Prepare(query)
	commentsTable.Exec()
	query = `CREATE TABLE IF NOT EXISTS interaction_comments (
		comment_id INTEGER NOT NULL,
		post_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		interaction TEXT NOT NULL,
		FOREIGN KEY("comment_id") REFERENCES comments("comment_id"),
		FOREIGN KEY("post_id") REFERENCES posts("post_id"),
		FOREIGN KEY("user_id") REFERENCES users("user_id")
	)`
	interaction_commentsTable, _ := db.Prepare(query)
	interaction_commentsTable.Exec()
	query = `CREATE TABLE IF NOT EXISTS requests (
		request_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		body TEXT NOT NULL,
		time_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY("request_id" AUTOINCREMENT),
		FOREIGN KEY("user_id") REFERENCES users("user_id")
	)`
	requestsTable, _ := db.Prepare(query)
	requestsTable.Exec()
	query = `CREATE TABLE IF NOT EXISTS actions (
		request_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		body TEXT NOT NULL,
		time_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY("request_id") REFERENCES requests("request_id"),
		FOREIGN KEY("user_id") REFERENCES users("user_id")
	)`
	actionsTable, _ := db.Prepare(query)
	actionsTable.Exec()
	row, _ := db.Query("SELECT user_id, user_name FROM users WHERE user_id = 1")
	var user_id int
	var user_name string
	for row.Next() {
		row.Scan(&user_id, &user_name)
		break
	}
	if user_id != 1 && user_name != "admin" {
		usersTable, _ = db.Prepare("INSERT INTO users (user_name, user_email, user_pass, user_type) VALUES (?, ?, ?, ?)")
		usersTable.Exec("admin", "admin", "admin", "admin")
	}
}
