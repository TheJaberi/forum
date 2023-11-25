package forum

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func CreateTables() {
	Database, err := sql.Open("sqlite3", "./forum.db") // The Database must be opened and assigned to a variable
	if err != nil {
		log.Fatal(err)
	}
	_, errpragma := Database.Exec("PRAGMA foreign_keys = ON") // not sure if we need this but foreign keys are disabled by default
	if errpragma != nil {
		log.Fatal(errpragma)
	}
	usertable, err2 := Database.Prepare("CREATE TABLE IF NOT EXISTS Users (id INTEGER PRIMARY KEY, username TEXT NOT NULL, Password TEXT NOT NULL, email TEXT)")
	if err2 != nil { // table for users is created if it does not exist in line 24
		log.Fatal(err2)
	}
	usertable.Exec() // Exec executes the query which was Prepared in line 20
	postsTable, err4 := Database.Prepare("CREATE TABLE IF NOT EXISTS Posts (id INTEGER PRIMARY KEY, Title TEXT, body TEXT, user_id INTEGER)")
	if err4 != nil { // table for posts is created if it doesnot exist
		log.Fatal(err4)
	}
	postsTable.Exec() // Exec executes query on line 25
	categoryTable, err5 := Database.Prepare("CREATE TABLE IF NOT EXISTS Category (id INTEGER PRIMARY KEY, Name TEXT)")
	if err5 != nil { // table for category is created if it doesnot exist
		log.Fatal(err5)
	}
	categoryTable.Exec() // Exec executes query on line 29
	post2categoryTable, err6 := Database.Prepare("CREATE TABLE IF NOT EXISTS Post2Category (id INTEGER PRIMARY KEY, post_id INTEGER NOT NULL, category_id INTEGER NOT NULL)")
	if err6 != nil { // table for linking posts and categories is created if it doesnot exist
		log.Fatal(err6)
	}
	post2categoryTable.Exec() // Exec executes query on line 34
	interactionTable, err7 := Database.Prepare("CREATE TABLE IF NOT EXISTS Interaction (id INTEGER PRIMARY KEY, post_id INTEGER NOT NULL, user_id INTEGER NOT NULL, interaction BIT NOT NULL)")
	if err7 != nil { // table for likes and dislikes is created if it doesnot exist
		log.Fatal(err7)
	}
	interactionTable.Exec()
	defer Database.Close()
}
