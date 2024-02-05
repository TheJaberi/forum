package forum

import (
	"context"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// TABLE: interaction_comments

func InsertCommentInteraction(post_id int, user_id int, likeOrDislike int, post bool, comment_id int) {
	query := "INSERT INTO `interaction_comments` (`comment_id`, `post_id`, `user_id`, `interaction`) VALUES (?, ?, ?, ?)"
	_, err := DB.ExecContext(context.Background(), query, comment_id, post_id, user_id, likeOrDislike)
	if err != nil { // the post is added using the ExecContext along with the userid which is in the LoggedUser variable
		log.Fatal(err)
	}
}
func RemoveCommentInteraction(post_id int, user_id int, post bool, comment_id int) {
	query := "DELETE FROM `interaction_comments` where comment_id = ? AND user_id = ?"
	_, err := DB.ExecContext(context.Background(), query, comment_id, user_id)
	if err != nil { // the post is added using the ExecContext along with the userid which is in the LoggedUser variable
		log.Fatal(err)
	}
}
func UpdateCommentInteraction(post_id int, user_id int, likeOrDislike int, post bool, comment_id int) {
	query := "UPDATE interaction_comments SET interaction = ? where comment_id= ? AND user_id = ?"
	_, err := DB.ExecContext(context.Background(), query, likeOrDislike, comment_id, user_id)
	if err != nil { // the post is added using the ExecContext along with the userid which is in the LoggedUser variable
		log.Fatal(err)
	}
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
			err = GetCommentLikes(&AllData.Postpage.Comments[i])
			if err != nil {
				return err
			}
			err = GetCommentDislikes(&AllData.Postpage.Comments[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func GetCommentLikes(c *Comment) error {
	likeCommentdata := DB.QueryRow("SELECT COUNT(user_id) FROM interaction_comments where comment_id = ? AND interaction = ?", c.Comment_id, 1) // to present the numb of likes for each comment
	err := likeCommentdata.Scan(&c.Likes)
	if err != nil {
		return errors.New("Post Interaction (Likes) Scan Error:" + err.Error())
	}
	return nil
}
func GetCommentDislikes(c *Comment) error {
	dislikedata := DB.QueryRow("SELECT COUNT(user_id) FROM interaction_comments where comment_id = ? AND interaction = ?", c.Comment_id, 0) // to present the numb of likes for each comment
	err := dislikedata.Scan(&c.Dislikes)
	if err != nil {
		return errors.New("Post Interaction (Dislikes) Scan Error:" + err.Error())
	}
	return nil
}
