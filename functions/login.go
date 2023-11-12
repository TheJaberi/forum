package forum

import (
	"fmt"
	"log"
	"database/sql"
	_"github.com/mattn/go-sqlite3"
)

func Login(username string, password string) {
	var userid int
	var usernametmp string
	var passwordtmp string
	var emailtmp string
	Database, errdatabase := sql.Open("sqlite3", "./forum.db")
	if errdatabase != nil {
		log.Fatal(errdatabase)
	}
	userdata, err := Database.Query("SELECT id, username, password, email FROM Users") // select gets the data from users table
	if err != nil {
		log.Fatal(err)
	}
	// for loop over all the rows until a match with the username is found
	for userdata.Next() { // Next prepares the next row in the table for reading using scan
		userdata.Scan(&userid, &usernametmp, &passwordtmp, &emailtmp) // scan assigns the data of the row to variables
		if username == usernametmp { // check if the username matches the username in the row
			if password == passwordtmp { // check if the password of the user is matches too
				LoggedUser.Userid = userid // if the password matches assign the data from the row to a global variable
				LoggedUser.Username = usernametmp
				LoggedUser.Password = passwordtmp
				LoggedUser.Email = emailtmp
				LoggedUser.Registered = true // if there is a match registered boolean assigned true
				defer userdata.Close()
			} else {
				ErrorMsg = "Wrong Password try again"
			}
		}
	}
	
	if LoggedUser.Username == "" {
		ErrorMsg = "Username entered is not registered"
	}
	fmt.Println(LoggedUser)
	defer Database.Close()
}
