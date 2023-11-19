package forum

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func ViewPosts() {
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
		likedata, likeerr := Database.Query("Select COUNT from Interaction where post_id = ? AND interaction = true", posttmp.PostID)
		if likeerr != nil {
			log.Fatal(likeerr)
		}
		likedata.Scan(&posttmp.Likes)
		dislikedata, dislikeerr := Database.Query("Select COUNT from Interaction where post_id = ? AND interaction = false", posttmp.PostID)
		if dislikeerr != nil {
			log.Fatal(dislikeerr)
		}
		dislikedata.Scan(&posttmp.Dislikes)
		AllPosts = append(AllPosts, posttmp)
	}
}
