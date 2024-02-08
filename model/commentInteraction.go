package forum

import (
	"context"
	"errors"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// TABLE: interaction_comments

func CommentInteraction(add, remove, path string) (Post, error) {
	var p Post
	var addComment_id int
	var remComment_id int
	var err error
	if add != "" {
		addComment_id, err = strconv.Atoi(add) // comment interaction handles the data from like or dislike button if the user logged hasn't already clicked on it
		if err != nil {
			return p, err
		}
	}
	if remove != "" {
		remComment_id, err = strconv.Atoi(remove) // remove interaction handles the data from like or dislike button if the user logged has already clicked on it
		if err != nil {
			return p, err
		}
	}
	user_id := LoggedUser.Userid
	commentPos := 0
	for i := 0; i < len(AllData.Postpage.Comments); i++ {
		if AllData.Postpage.Comments[i].Comment_id == addComment_id || AllData.Postpage.Comments[i].Comment_id == remComment_id {
			commentPos = i
			break
		}
	}
	if addComment_id > remComment_id { // which ever value is greater determines whether to add or remove
		if path == "/commentlike/" {
			if !AllData.Postpage.Comments[commentPos].CommentUserDislike {
				err := InsertCommentInteraction(AllData.Postpage.PostID, user_id, 1, addComment_id) // insert adds the interaction to the database 1 is like 0 is dislike
				if err != nil {
					return p, err
				}
				AllData.Postpage.Comments[commentPos].CommentUserlike = true // changes the comment like or dislike for the logged in user in the all comments var
			} else {
				err := UpdateCommentInteraction(user_id, 1, addComment_id) // update is used if a like has to be changed to a dislike or vice versa
				if err != nil {
					return p, err
				}
				AllData.Postpage.Comments[commentPos].CommentUserlike = true
				AllData.Postpage.Comments[commentPos].CommentUserDislike = false
			}
		} else {
			if !AllData.Postpage.Comments[commentPos].CommentUserlike {
				err := InsertCommentInteraction(AllData.Postpage.PostID, user_id, 0, addComment_id)
				if err != nil {
					return p, err
				}
				AllData.Postpage.Comments[commentPos].CommentUserDislike = true
			} else {
				err := UpdateCommentInteraction(user_id, 0, addComment_id)
				if err != nil {
					return p, err
				}
				AllData.Postpage.Comments[commentPos].CommentUserDislike = true
				AllData.Postpage.Comments[commentPos].CommentUserlike = false
			}
		}
	} else {
		err := RemoveCommentInteraction(user_id, remComment_id) // remove is greater means there is already an interaction that needs to be removed
		if err != nil {
			return p, err
		}
		AllData.Postpage.Comments[commentPos].CommentUserlike = false
		AllData.Postpage.Comments[commentPos].CommentUserDislike = false
	}
	AllData.Postpage.LoggedUser = true
	return AllData.Postpage, nil
}
func InsertCommentInteraction(postID int, userID int, likeOrDislike int, commentID int) error {
	query := "INSERT INTO `interaction_comments` (`comment_id`, `post_id`, `user_id`, `interaction`) VALUES (?, ?, ?, ?)"
	_, err := DB.ExecContext(context.Background(), query, commentID, postID, userID, likeOrDislike)
	if err != nil { // the post is added using the ExecContext along with the userid which is in the LoggedUser variable
		log.Println(err)
		return err
	}
	return nil
}
func RemoveCommentInteraction(userID int, commentID int) error {
	query := "DELETE FROM `interaction_comments` where comment_id = ? AND user_id = ?"
	_, err := DB.ExecContext(context.Background(), query, commentID, userID)
	if err != nil { // the post is added using the ExecContext along with the userid which is in the LoggedUser variable
		log.Println(err)
		return err
	}
	return nil
}
func UpdateCommentInteraction(userID int, likeOrDislike int, commentID int) error {
	query := "UPDATE interaction_comments SET interaction = ? where comment_id= ? AND user_id = ?"
	_, err := DB.ExecContext(context.Background(), query, likeOrDislike, commentID, userID)
	if err != nil { // the post is added using the ExecContext along with the userid which is in the LoggedUser variable
		log.Println(err)
		return err
	}
	return nil
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
