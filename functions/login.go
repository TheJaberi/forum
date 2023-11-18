package forum

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func Login(username string, password string) {
	Database, errdatabase := sql.Open("sqlite3", "./forum.db")
	if errdatabase != nil {
		log.Fatal(errdatabase)
	}
	userdata := Database.QueryRow("SELECT id, username, password, email FROM Users where username = ? AND password = ?", username, password) // select gets the data from users table
	err := userdata.Scan(&LoggedUser.Userid, &LoggedUser.Username, &LoggedUser.Password, &LoggedUser.Email)                                  // scan assigns the data of the row to variables
	if err != nil {
		log.Fatal(err)
	} else {
		LoggedUser.Registered = true
	}
	defer Database.Close()
}
