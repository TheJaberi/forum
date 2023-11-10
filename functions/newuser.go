package forum

import ("database/sql"
"log"
_"github.com/mattn/go-sqlite3")
func NewUser(username string, password string){
	Database, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	usersTable, errprepare := Database.Prepare("INSERT INTO Users (username, Password) VALUES (?, ?)")
	if errprepare != nil {// preparing to insert the username and password data for line 15
		log.Fatal(errprepare)
	}
	usersTable.Exec(username, password) // Exec executes the query prepared on line 11 
	defer Database.Close()
}