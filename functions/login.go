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
	Database, _ := sql.Open("sqlite3", "./forum.db")
	userdata, err := Database.Query("SELECT id, username, password, email FROM Users")
	if err != nil {
		log.Fatal(err)
	}
	for userdata.Next() {
		fmt.Println(22)
		userdata.Scan(&userid, &usernametmp, &passwordtmp, &emailtmp)
		if username == usernametmp {
			if password == passwordtmp {
				LoggedUser.Userid = userid
				LoggedUser.Username = usernametmp
				LoggedUser.Password = passwordtmp
				LoggedUser.Email = emailtmp
				LoggedUser.Registered = true
				break
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
