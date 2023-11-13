package forum

import ("database/sql"
"log"
_"github.com/mattn/go-sqlite3")
func NewUser(username string, password string, email string){
	Database, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	usersTable, errprepare := Database.Prepare("INSERT INTO Users (username, Password, email) VALUES (?, ?, ?)")
	if errprepare != nil {// preparing to insert the username and password data for line 15
		log.Fatal(errprepare)
	}
	usersTable.Exec(username, password, email) // Exec executes the query prepared on line 11 
	defer Database.Close()
}
