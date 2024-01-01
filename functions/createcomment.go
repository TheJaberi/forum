package forum

import (
	"context"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func CreateComment(commentContent string, postID int) {
	query := "INSERT INTO `comments` (`post_id`, `user_id`, `body`) VALUES (?, ?, ?)"
	_, err2 := DB.ExecContext(context.Background(), query, postID, LoggedUser.Userid, commentContent)
	if err2 != nil {
		log.Fatal(err2)
	}
	var comment Comment
	comment.Body = commentContent
	comment.post_id = postID
	comment.user_id = LoggedUser.Userid
	comment.username = LoggedUser.Username
	AllData.AllPosts[postID].Comments = append(AllData.AllPosts[postID].Comments, comment)
}
