package forum

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func CreatePost(title string, body string) {
	fmt.Println(LoggedUser)
	if LoggedUser.Registered{
		Database, err := sql.Open("sqlite3", "./forum.db")
		if err != nil{
			log.Fatal(err)
		}
	postsTable, err2 := Database.Prepare("INSERT INTO Posts (Title, body) VALUES (?, ?)")
	postsTable.Exec(title, body)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(title)
	fmt.Println(body)
	fmt.Println(LoggedUser.Userid)
}else {
	ErrorMsg = "Cannot create post need to log in first"
	fmt.Println(ErrorMsg)
}
}
