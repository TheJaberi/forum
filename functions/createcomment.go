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
	var commenTmp Comment
	commenTmp.Post_id = postID
	commenTmp.User_id = AllData.LoggedUser.Userid
	commenTmp.CommentUsername = AllData.LoggedUser.Username
	commenTmp.Body = commentContent
	AllData.AllPosts[postID-1].Comments = append(AllData.AllPosts[postID-1].Comments, commenTmp)
}
