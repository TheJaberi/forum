package forum

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ViewPosts() {
	AllPosts = nil
	Database, errdatabase := sql.Open("sqlite3", "./forum.db")
	if errdatabase != nil {
		log.Fatal(errdatabase)
	}
	defer Database.Close()
	postData, errpost := Database.Query("Select * from Posts")
	if errpost != nil {
		log.Fatal(errpost)
	}
	for postData.Next() {
		var posttmp Post
		postData.Scan(&posttmp.PostID, &posttmp.Title, &posttmp.Body, &posttmp.UserID)
		userData := Database.QueryRow("Select username from Users where id = ?", posttmp.UserID)
		userData.Scan(&posttmp.Username)
		categorydata, categoryerr := Database.Query("Select category_id from Post2Category where post_id = ?", posttmp.PostID)
		if categoryerr != nil {
			log.Fatal(categoryerr)
		}
		for categorydata.Next() {
			var categorytmp int
			categorydata.Scan(&categorytmp)
			for i := 0; i < len(AllCategories); i++ {
				if categorytmp == AllCategories[i].CategoryID {
					posttmp.Category = append(posttmp.Category, AllCategories[i])
					break
				}
			}
		}
		likedata := Database.QueryRow("SELECT COUNT(user_id) FROM Interaction where post_id = ? AND interaction = ?", posttmp.PostID, 1)
		likedata.Scan(&posttmp.Likes)
		dislikedata := Database.QueryRow("SELECT COUNT(user_id) FROM Interaction where post_id = ? AND interaction = ?", posttmp.PostID, 0)
		dislikedata.Scan(&posttmp.Dislikes)
		AllPosts = append(AllPosts, posttmp)
	}
	if LoggedUser.Registered{
	for i:= 0;i<len(AllPosts);i++{
		var interaction int
		postData := Database.QueryRow("SELECT interaction from Interaction where post_id = ? AND user_id = ?", i+1, LoggedUser.Userid)
		errpost := postData.Scan(&interaction)
		if errpost!=nil{
			fmt.Println(errpost)
			continue
		} else {
			if interaction==1{
				AllPosts[i].Userlike = true
			} else {
				AllPosts[i].UserDislike = true
			}
		}
	}}
}
