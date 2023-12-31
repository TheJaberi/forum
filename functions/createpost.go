package forum

import (
	"context"
	// "database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func CreatePost(title string, body string, postCategories []int) {
	var postData Post
	postData.Title = title
	postData.Body = body
	postData.UserID = 1
	postData.Username = "test"
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
		AllPosts = append(AllPosts, postData)
  		ViewPosts()	}