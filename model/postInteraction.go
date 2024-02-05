package forum

import (
	"context"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// TABLE: Interactions

func InsertPostInteraction(post_id int, user_id int, likeOrDislike int, post bool, comment_id int) {
	query := "INSERT INTO `Interaction` (`post_id`, `user_id`, `interaction`) VALUES (?, ?, ?)"
	_, err := DB.ExecContext(context.Background(), query, post_id, user_id, likeOrDislike)
	if err != nil { // the post is added using the ExecContext along with the userid which is in the LoggedUser variable
		log.Fatal(err)
	}
}

func RemovePostInteraction(post_id int, user_id int, post bool, comment_id int) {
	query := "DELETE FROM `Interaction` where post_id = ? AND user_id = ?"
	_, err := DB.ExecContext(context.Background(), query, post_id, user_id)
	if err != nil { // the post is added using the ExecContext along with the userid which is in the LoggedUser variable
		log.Fatal(err)
	}
}

func UpdatePostInteraction(post_id int, user_id int, likeOrDislike int, post bool, comment_id int) {
	query := "UPDATE Interaction SET interaction = ? where post_id= ? AND user_id = ?"
	_, err := DB.ExecContext(context.Background(), query, likeOrDislike, post_id, user_id)
	if err != nil { // the post is added using the ExecContext along with the userid which is in the LoggedUser variable
		log.Fatal(err)
	}
}

func GetUserPostInteractions() error {
	if LoggedUser.Registered { // FIXME Possibly redundent
		for i := range AllPosts {
			var interaction int
			postData := DB.QueryRow("SELECT interaction from Interaction where post_id = ? AND user_id = ?", i+1, LoggedUser.Userid)
			errpost := postData.Scan(&interaction)
			if errpost != nil {
				// FIXME what is the case for this?
				// continue
				return errors.New("Error scanning user post interactions")
			}
			if interaction == 1 {
				AllPosts[i].Userlike = true
			} else {
				AllPosts[i].UserDislike = true
			}
		}
	} else {
		return errors.New("User must log in to get post interactions")
	}
	return nil
}

func GetPostLikes(p *Post) error {
	likedata := DB.QueryRow("SELECT COUNT(user_id) FROM Interaction where post_id = ? AND interaction = ?", p.PostID, 1)
	err := likedata.Scan(&p.Likes)
	if err != nil {
		return errors.New("Post Interaction (Likes) Scan Error:" + err.Error())
	}
	return nil
}

func GetPostDislikes(p *Post) error {
	dislikedata := DB.QueryRow("SELECT COUNT(user_id) FROM Interaction where post_id = ? AND interaction = ?", p.PostID, 0)
	err := dislikedata.Scan(&p.Dislikes)
	if err != nil {
		return errors.New("Post Interaction (Dislikes) Scan Error:" + err.Error())
	}
	return nil
}
