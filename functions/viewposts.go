package forum

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func ViewPosts(){
	var count int
	Database, errdatabase := sql.Open("sqlite3", "./forum.db")
	if errdatabase != nil {
		log.Fatal(errdatabase)
	}
	defer Database.Close()
	countRows:= Database.QueryRow("SELECT COUNT(*) FROM Posts")
	countRows.Scan(&count)
	// fmt.Println(count)
	for i:=1;i<=count;i++{
	var posttmp Post
	postData := Database.QueryRow("Select * from Posts where id = ?", i)
	postData.Scan(&posttmp.PostID, &posttmp.Title, &posttmp.Body, &posttmp.UserID)
		// if errPost != nil {
		// 	log.Fatal(errPost)
		// }
		userData := Database.QueryRow("Select username from Users where id = ?", posttmp.UserID)
		userData.Scan(&posttmp.Username)
			// if errUser != nil {
			// 	log.Fatal(errUser)
			// }
		if i<len(AllCategories){
		posttmp.Categories = AllCategories[i]}
		AllPosts = append(AllPosts, posttmp)
	}
	for i:= count;i<len(AllCategories);i++{
		var posttmp Post
		posttmp.Categories = AllCategories[i]
		AllPosts = append(AllPosts, posttmp)
	}
	}
