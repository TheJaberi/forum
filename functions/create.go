package forum

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func CreatePost(title string, body string, postCategories []int) {
	var postData Post
	postData.Title = title
	postData.Body = body
	postData.UserID = LoggedUser.Userid
	postData.Username = LoggedUser.Username
	// if LoggedUser.Registered { // check if registered is to true to add the post to the database
		query := "INSERT INTO `posts` (`Title`, `body`, `user_id`) VALUES (?, ?, ?)"
		rowdata, err2 := DB.ExecContext(context.Background(), query, title, body, LoggedUser.Userid)
		if err2 != nil { // the post is added using the ExecContext along with the userid which is in the LoggedUser variable
			log.Fatal(err2)
		}
		postid, err := rowdata.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < len(postCategories); i++ {
			queryCategory := "INSERT INTO `Post2Category` (`post_id`, `category_id`) VALUES (?, ?)"
			_, err3 := DB.ExecContext(context.Background(), queryCategory, postid, postCategories[i])
			if err3 != nil { // the post is added using the ExecContext along with the userid which is in the LoggedUser variable
				log.Fatal(err3)
			}
			for j := 0; j < len(AllCategories); j++ {
				if AllCategories[j].CategoryID == postCategories[i] {
					postData.Category = append(postData.Category, AllCategories[j])
					break
				}
			}
		}
		postData.TimeCreated = "Now"
		AllPosts = append(AllPosts, postData)
		AllData.AllPosts = AllPosts
}
func CreateComment(commentContent string, postID int) {
	query := "INSERT INTO `comments` (`post_id`, `user_id`, `body`) VALUES (?, ?, ?)"
	_, err2 := DB.ExecContext(context.Background(), query, postID, LoggedUser.Userid, commentContent)
	if err2 != nil {
		log.Fatal(err2)
	}
	var commenTmp Comment
	commenTmp.Post_id = postID
	commenTmp.User_id = AllData.LoggedUser.Userid
	commenTmp.CommentUsername = AllData.LoggedUser.Username
	commenTmp.Body = commentContent
	AllData.AllPosts[postID-1].Comments = append(AllData.AllPosts[postID-1].Comments, commenTmp)
}
func CreateCategory(name string) {
	var cat Category
	query := "INSERT INTO `Category` (`Name`) VALUES (?)"
	_, err2 := DB.ExecContext(context.Background(), query, name)
	if err2 != nil { // the category is added using the ExecContext
		log.Fatal(err2)
	}
	cat.CategoryName = name
	AllCategories = append(AllCategories, cat)
}