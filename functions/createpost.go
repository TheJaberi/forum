package forum

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CreatePost(title string, body string) {
	fmt.Println(LoggedUser)
	if LoggedUser.Registered{  // check if registered is to true to add the post to the database
		Database, err := sql.Open("sqlite3", "./forum.db")
		if err != nil{
			log.Fatal(err)
		}
	query := "INSERT INTO `Posts` (`Title`, `body`, `user_id`) VALUES (?, ?, ?)" 
	_, err2 := Database.ExecContext(context.Background(),query, title, body, LoggedUser.Userid)
	if err2 != nil { // the post is added using the ExecContext along with the userid which is in the LoggedUser variable
		log.Fatal(err2)
	}
}else {
	ErrorMsg = "Cannot create post need to log in first"
	fmt.Println(ErrorMsg)
}
defer Database.Close()
}
