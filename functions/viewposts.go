package forum

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func ViewPosts(){
	// var count int
	Database, errdatabase := sql.Open("sqlite3", "./forum.db")
	if errdatabase != nil {
		log.Fatal(errdatabase)
	}
	defer Database.Close()
	// countRows:= Database.QueryRow("SELECT COUNT(*) FROM Posts")
	// countRows.Scan(&count)
	// fmt.Println(count)
	postData, errpost := Database.Query("Select * from Posts")
	if errpost != nil{
		log.Fatal(errpost)
	}
	for postData.Next(){
		var posttmp Post
	postData.Scan(&posttmp.PostID, &posttmp.Title, &posttmp.Body, &posttmp.UserID)
		// if errPost != nil {
		// 	log.Fatal(errPost)
		// }
		userData := Database.QueryRow("Select username from Users where id = ?", posttmp.UserID)
		userData.Scan(&posttmp.Username)
			// if errUser != nil {
			// 	log.Fatal(errUser)
			// }
		categorydata, categoryerr := Database.Query("Select category_id from Post2Category where post_id = ?", posttmp.PostID)

		if categoryerr != nil {
			log.Fatal(categoryerr)
		}
		for categorydata.Next(){
			var categorytmp int
			categorydata.Scan(&categorytmp)
			for i:=0;i<len(AllCategories);i++{
				if categorytmp == AllCategories[i].CategoryID {
					posttmp.Category = append(posttmp.Category, AllCategories[i])
					break
				}
			}
		}
		AllPosts = append(AllPosts, posttmp)
	}
	}
