package forum

import ("database/sql"
"log"
_"github.com/mattn/go-sqlite3")
func NewUser(username string, password string){
	Database, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	usersTable, _ := Database.Prepare("INSERT INTO Users (username, Password) VALUES (?, ?)")
	usersTable.Exec(username, password)
}