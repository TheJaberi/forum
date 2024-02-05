package forum

import (
	"context"
	"errors"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// TABLE: Interactions

func PostInteractions(add, remove, path string) (Post, error) {
	var p Post
	addPost_id, err := strconv.Atoi(add) // post interaction handles the data from like or dislike button if the user logged hasn't already clicked on it
	if err != nil {
		return p, err
	}
	remPost_id, err := strconv.Atoi(remove) // remove interaction handles the data from like or dislike button if the user logged has already clicked on it
	if err != nil {
		return p, err
	}
	user_id := LoggedUser.Userid
	if addPost_id > remPost_id { // which ever value is greater determines whether to add or remove
		p = AllPosts[addPost_id-1]
		if path == "/like/" {
			if !AllPosts[addPost_id-1].UserDislike {
				InsertPostInteraction(addPost_id, user_id, 1, true, 0) // insert adds the interaction to the database 1 is like 0 is dislike
				AllPosts[addPost_id-1].Userlike = true                 // changes the post like or dislike for the logged in user in the all posts var
			} else {
				UpdatePostInteraction(addPost_id, user_id, 1, true, 0) // update is used if a like has to be changed to a dislike or vice versa
				AllPosts[addPost_id-1].Userlike = true
				AllPosts[addPost_id-1].UserDislike = false
			}
		} else {
			if !AllPosts[addPost_id-1].Userlike {
				InsertPostInteraction(addPost_id, user_id, 0, true, 0)
				AllPosts[addPost_id-1].UserDislike = true
			} else {
				UpdatePostInteraction(addPost_id, user_id, 0, true, 0)
				AllPosts[addPost_id-1].UserDislike = true
				AllPosts[addPost_id-1].Userlike = false
			}
		}
		p = AllPosts[addPost_id-1]
	} else {
		RemovePostInteraction(remPost_id, user_id, true, 0) //remove is greater means there is already an interaction that needs to be removed
		AllPosts[remPost_id-1].Userlike = false
		AllPosts[remPost_id-1].UserDislike = false
		p = AllPosts[remPost_id-1]
	}
	p.LoggedUser = true
	return p, nil
}

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
			err := postData.Scan(&interaction)
			if err == nil {
				if interaction == 1 {
					AllPosts[i].Userlike = true
				} else {
					AllPosts[i].UserDislike = true
				}
			} else {
				log.Printf(err.Error())
				continue // used for logout (remove user post interactions from global struct)?
			}
		}
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
