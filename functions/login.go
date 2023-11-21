package forum

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func Login(username string, password string) {
	Database, errdatabase := sql.Open("sqlite3", "./forum.db")
	if errdatabase != nil {
		log.Fatal(errdatabase)
	}
	userdata := Database.QueryRow("SELECT id, username, password, email FROM Users where username = ? AND password = ?", username, password) // select gets the data from users table
	err := userdata.Scan(&LoggedUser.Userid, &LoggedUser.Username, &LoggedUser.Password, &LoggedUser.Email)                                  // scan assigns the data of the row to variables
	if err != nil {
		fmt.Println(err)
			} else {
		LoggedUser.Registered = true
	}
	for i:= 0;i<len(AllPosts);i++{
		var interaction int
		postData := Database.QueryRow("SELECT interaction where post_id = ?, user_id = ?", i+1, LoggedUser.Userid)
		errpost := postData.Scan(&interaction)
		if errpost!=nil{
			continue
		} else {
			if interaction==1{
				AllPosts[i].Userlike = true
			} else {
				AllPosts[i].UserDislike = true
			}
		}
	}
	AllData.LoggedUser = LoggedUser
	AllData.LoggedUserID = LoggedUser.Userid
	AllData.IsLogged = true
	ViewPosts()
	defer Database.Close()
}
