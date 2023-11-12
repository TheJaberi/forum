package forum 

import (
	// "fmt"
	"log"
	"database/sql"
	_"github.com/mattn/go-sqlite3"
)

func ViewPosts(){
	// var post Post
	var posttmp Post
	var userIDtmp int
	Database, errdatabase := sql.Open("sqlite3", "./forum.db")
	if errdatabase != nil {
		log.Fatal(errdatabase)
	}
	defer Database.Close()
	postData, errpost := Database.Query("Select id, Title, Body, user_id from Posts")
	if errpost != nil {
		log.Fatal(errpost)
	}
	defer postData.Close()
	for postData.Next(){
		postData.Scan(&posttmp.PostID, &posttmp.Title, &posttmp.Body, &posttmp.UserID)
		userData, usererr := Database.Query("Select id, username from Users")
		if usererr != nil {
			log.Fatal(usererr)
		}
		for userData.Next(){
			userData.Scan(&userIDtmp, &posttmp.Username)
			if posttmp.UserID == userIDtmp {
				userData.Close()
			}
		}
		AllPosts = append(AllPosts, posttmp)
	}
	}
