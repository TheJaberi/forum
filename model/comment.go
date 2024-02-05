package forum

import (
	"context"
	"errors"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func CreateComment(commentContent string, postID int) {
	var c = Comment{
		Post_id:         postID,
		User_id:         AllData.LoggedUser.Userid,
		CommentUsername: AllData.LoggedUser.Username,
		Body:            commentContent,
	}
	id, err := createCommentDb(c)
	if err != nil {
		// TODO RETURN ERROR
	}
	c, err = getComment(id)
	AllData.AllPosts[c.Post_id-1].Comments = append(AllData.AllPosts[c.Post_id-1].Comments, c)
}
func GetUserCommentInteractions() error {
	if LoggedUser.Registered { // if the user is logged in the fact that he has liked or disliked the post is saved in all posts
		for i := 0; i < len(AllData.Postpage.Comments); i++ {
			var interaction int
			postData := DB.QueryRow("SELECT interaction from interaction_comments where comment_id = ? AND user_id = ?", AllData.Postpage.Comments[i].Comment_id, LoggedUser.Userid)
			err := postData.Scan(&interaction)
			if err != nil {
				continue
			} else {
				if interaction == 1 {
					AllData.Postpage.Comments[i].CommentUserlike = true
				} else {
					AllData.Postpage.Comments[i].CommentUserDislike = true
				}
			}
			err = getCommentLikes(&AllData.Postpage.Comments[i])
			if err != nil {
				return err
			}
			err = getCommentDislikes(&AllData.Postpage.Comments[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func createCommentDb(c Comment) (int64, error) {
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

func getComment(id int64) (Comment, error) {
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
func GetCommentUsername(c *Comment) error {
	userData := DB.QueryRow("Select user_name from users where user_id = ?", c.User_id)
	err := userData.Scan(&c.CommentUsername)
	if err != nil {
		return errors.New("Comment Username Scan Error:" + err.Error())
	}
	return nil
}
func getCommentLikes(c *Comment) error {
	likeCommentdata := DB.QueryRow("SELECT COUNT(user_id) FROM interaction_comments where comment_id = ? AND interaction = ?", c.Comment_id, 1) // to present the numb of likes for each comment
	err := likeCommentdata.Scan(&c.Likes)
	if err != nil {
		return errors.New("Post Interaction (Likes) Scan Error:" + err.Error())
	}
	return nil
}

func getCommentDislikes(c *Comment) error {
	dislikedata := DB.QueryRow("SELECT COUNT(user_id) FROM interaction_comments where comment_id = ? AND interaction = ?", c.Comment_id, 0) // to present the numb of likes for each comment
	err := dislikedata.Scan(&c.Dislikes)
	if err != nil {
		return errors.New("Post Interaction (Dislikes) Scan Error:" + err.Error())
	}
	return nil
}
