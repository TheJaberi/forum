package forum 

import (
	// "fmt"
	"log"
	"database/sql"
	_"github.com/mattn/go-sqlite3"
)

func ViewPosts(){
	var posttmp Post
	Database, errdatabase := sql.Open("sqlite3", "./forum.db")
	if errdatabase != nil {
		log.Fatal(errdatabase)
	}
	defer Database.Close()
	postData := Database.QueryRow("Select * from Posts")
	postData.Scan(&posttmp.PostID, &posttmp.Title, &posttmp.Body, &posttmp.UserID)
		// if errPost != nil {
		// 	log.Fatal(errPost)
		// }
		userData := Database.QueryRow("Select username from Users where user_id = ?", posttmp.UserID)
		userData.Scan(&posttmp.Username)
			// if errUser != nil {
			// 	log.Fatal(errUser)
			// }
		AllPosts = append(AllPosts, posttmp)
	}
