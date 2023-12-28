package forum

import (
	// "Database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

func ViewPosts() {
	AllPosts = nil
	// Database, errdatabase := sql.Open("sqlite3", "./forum.db")
	// if errdatabase != nil {
	// 	log.Fatal(errdatabase)
	// }
	// defer Database.Close()
	postData, errpost := DB.Query("Select id, Title, body, user_id from posts")
	if errpost != nil {
		log.Fatal(errpost)
	}
	for postData.Next() { // this loop ends at the end of the function since it needs to get the data for each post from 5 tables
		var posttmp Post // temporary type post is appended to all posts at the end of the loop after gathering all of the data
		postData.Scan(&posttmp.PostID, &posttmp.Title, &posttmp.Body, &posttmp.UserID) // posts id, title, body and user id is from posts table
		userData := DB.QueryRow("Select user_name from users where id = ?", posttmp.UserID) // username of the user who posted
		userData.Scan(&posttmp.Username)
		categorydata, categoryerr := DB.Query("Select category_id from Post2Category where post_id = ?", posttmp.PostID) // link between posts and its categories
		if categoryerr != nil {
			log.Fatal(categoryerr)
		}
		for categorydata.Next() {
			var categorytmp int
			categorydata.Scan(&categorytmp)
			for i := 0; i < len(AllCategories); i++ { // the name of the categories is already saved in all categories 
				if categorytmp == AllCategories[i].CategoryID {
					posttmp.Category = append(posttmp.Category, AllCategories[i])
					break
				}
			}
		}
		likedata := DB.QueryRow("SELECT COUNT(user_id) FROM Interaction where post_id = ? AND interaction = ?", posttmp.PostID, 1) // to present the numb of likes for each post
		likedata.Scan(&posttmp.Likes)
		dislikedata := DB.QueryRow("SELECT COUNT(user_id) FROM Interaction where post_id = ? AND interaction = ?", posttmp.PostID, 0) // to present the numb of dislikes for each post
		dislikedata.Scan(&posttmp.Dislikes)
		AllPosts = append(AllPosts, posttmp)
	}
	fmt.Println(AllPosts)
UpdatePosts()
}
