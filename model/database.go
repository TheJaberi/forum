package forum

import (
	"database/sql"
)

// creates all the tables in the database
func DataBase(db *sql.DB) {
	var query string
	query = `CREATE TABLE IF NOT EXISTS users (
		user_id INTEGER NOT NULL,
		user_name CHAR(10) NOT NULL UNIQUE,
		user_email CHAR(25) NOT NULL UNIQUE,
		user_pass PASSWORD NOT NULL,
		user_type TEXT NOT NULL DEFAULT member,
		time_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY("user_id" AUTOINCREMENT)
	);`
	usersTable, _ := db.Prepare(query)
	usersTable.Exec()
	query = `CREATE TABLE IF NOT EXISTS posts (
		id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		Title TEXT NOT NULL,
		img_url TEXT,
		body TEXT NOT NULL,
		time_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY("id" AUTOINCREMENT),
		FOREIGN KEY("user_id") REFERENCES users("user_id")
	);`
	postsTable, _ := db.Prepare(query)
	postsTable.Exec()
	query = `CREATE TABLE IF NOT EXISTS Interaction (
		id INTEGER PRIMARY KEY,
		post_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		interaction BIT NOT NULL);`
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
	);`
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
	query = `CREATE TABLE IF NOT EXISTS Category (
		id INTEGER PRIMARY KEY,
		Name TEXT
	)`
	categoryTable, _ := db.Prepare(query)
	categoryTable.Exec()
	query = `CREATE TABLE IF NOT EXISTS Post2Category
		(id INTEGER PRIMARY KEY,
		post_id INTEGER NOT NULL,
		category_id INTEGER NOT NULL
	)`
	post2categoryTable, _ := db.Prepare(query)
	post2categoryTable.Exec()
}
