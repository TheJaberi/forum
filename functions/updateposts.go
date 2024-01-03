package forum

import (
	_ "github.com/mattn/go-sqlite3"
)

func UpdatePosts() {
	if LoggedUser.Registered { // if the user is logged in the fact that he has liked or disliked the post is saved in all posts
		for i := 0; i < len(AllPosts); i++ {
			var interaction int
			postData := DB.QueryRow("SELECT interaction from Interaction where post_id = ? AND user_id = ?", i+1, LoggedUser.Userid)
			errpost := postData.Scan(&interaction)
			if errpost != nil {
				continue
			} else {
				if interaction == 1 {
					AllPosts[i].Userlike = true
				} else {
					AllPosts[i].UserDislike = true
				}
			}
		}
	}
}
