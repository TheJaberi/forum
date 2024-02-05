package forum

import (
	"context"
	"errors"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// TABLE: comments

func CreateComment(commentContent string, postID int) {
	var c = Comment{
		Post_id:         postID,
		User_id:         AllData.LoggedUser.Userid,
		CommentUsername: AllData.LoggedUser.Username,
		Body:            commentContent,
	}
	id, err := CreateCommentDb(c)
	if err != nil {
		// TODO RETURN ERROR
	}
	c, err = GetComment(id)
	AllData.AllPosts[c.Post_id-1].Comments = append(AllData.AllPosts[c.Post_id-1].Comments, c)
}

func CreateCommentDb(c Comment) (int64, error) {
	query := "INSERT INTO `comments` (`post_id`, `user_id`, `body`) VALUES (?, ?, ?)"
	rowData, err := DB.ExecContext(context.Background(), query, c.Post_id, c.User_id, c.Body)
	if err != nil {
		log.Fatal(err)
	}
	commentID, err := rowData.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return commentID, nil
}
func GetComment(id int64) (Comment, error) {
	row := DB.QueryRow("SELECT comment_id, post_id, user_id, body, time_created from comments WHERE id=?", id)
	var c Comment
	err := row.Scan(&c.Comment_id, &c.Post_id, &c.User_id, &c.Body, &c.TimeCreated)
	if err != nil {
		log.Printf("Error Getting Post")
		return c, err
	}
	c.TimeCreated = strings.Replace(c.TimeCreated, "T", " ", -1)
	c.TimeCreated = strings.Replace(c.TimeCreated, "Z", " ", -1)
	err = GetCommentUsername(&c)
	if err != nil {
		log.Println(err.Error())
		return c, err
	}
	return c, nil
}

func GetPostComments(p *Post) error {
	commentData, err := DB.Query("Select comment_id, body, user_id, time_created from comments where post_id = ?", p.PostID)
	if err != nil {
		return errors.New("Comment Query Error:" + err.Error())
	}
	defer commentData.Close()
	for commentData.Next() {
		var c Comment
		err := commentData.Scan(&c.Comment_id, &c.Body, &c.User_id, &c.TimeCreated)
		if err != nil {
			return errors.New("Comment Scan Error:" + err.Error())
		}
		c.TimeCreated = strings.Replace(c.TimeCreated, "T", " ", -1)
		c.TimeCreated = strings.Replace(c.TimeCreated, "Z", " ", -1)
		err = GetCommentUsername(&c)
		if err != nil {
			return err
		}
		p.Comments = append(p.Comments, c)
	}
	p.NumbOfComments = len(p.Comments)
	return nil
}
