package forum

import (
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func ViewPosts() {
	AllPosts = nil
	postData, errpost := DB.Query("Select id, Title, body, user_id, time_created from posts")
	if errpost != nil {
		log.Fatal(errpost)
	}
	for postData.Next() { // this loop ends at the end of the function since it needs to get the data for each post from 5 tables
		var posttmp Post                                                                                     // temporary type post is appended to all posts at the end of the loop after gathering all of the data
		postData.Scan(&posttmp.PostID, &posttmp.Title, &posttmp.Body, &posttmp.UserID, &posttmp.TimeCreated) // posts id, title, body and user id is from posts table
		posttmp.TimeCreated = strings.Replace(posttmp.TimeCreated, "T", " ", -1)
		posttmp.TimeCreated = strings.Replace(posttmp.TimeCreated, "Z", " ", -1)
		userData := DB.QueryRow("Select user_name from users where user_id = ?", posttmp.UserID) // username of the user who posted
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
		commentData, commenterr := DB.Query("Select body, user_id from comments where post_id = ?", posttmp.PostID) // link between posts and its categories
		if commenterr != nil {
			log.Fatal(categoryerr)
		}
		for commentData.Next() {
			var commenttmp Comment
			commentData.Scan(&commenttmp.Body, &commenttmp.User_id)
			userData := DB.QueryRow("Select user_name from users where user_id = ?", commenttmp.User_id)
			userData.Scan(&commenttmp.CommentUsername)
			posttmp.Comments = append(posttmp.Comments, commenttmp)
		}
		likedata := DB.QueryRow("SELECT COUNT(user_id) FROM Interaction where post_id = ? AND interaction = ?", posttmp.PostID, 1) // to present the numb of likes for each post
		likedata.Scan(&posttmp.Likes)
		dislikedata := DB.QueryRow("SELECT COUNT(user_id) FROM Interaction where post_id = ? AND interaction = ?", posttmp.PostID, 0) // to present the numb of dislikes for each post
		dislikedata.Scan(&posttmp.Dislikes)
		AllPosts = append(AllPosts, posttmp)
	}

	UpdatePosts()
	AllData.AllPosts = AllPosts
	// fmt.Println(AllData.AllPosts)
}
