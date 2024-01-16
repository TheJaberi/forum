package forum

import (
	_ "github.com/mattn/go-sqlite3"
)

func UpdateComments() {
	if LoggedUser.Registered { // if the user is logged in the fact that he has liked or disliked the post is saved in all posts
		for i := 0; i < len(AllData.Postpage.Comments); i++ {
			var interaction int
			postData := DB.QueryRow("SELECT interaction from interaction_comments where comment_id = ? AND user_id = ?", AllData.Postpage.Comments[i].Comment_id, LoggedUser.Userid)
			errpost := postData.Scan(&interaction)
			if errpost != nil {
				continue
			} else {
				if interaction == 1 {
					AllData.Postpage.Comments[i].CommentUserlike = true
				} else {
					AllData.Postpage.Comments[i].CommentUserDislike = true
				}
			}
			likeCommentdata := DB.QueryRow("SELECT COUNT(user_id) FROM interaction_comments where comment_id = ? AND interaction = ?", AllData.Postpage.Comments[i].Comment_id, 1) // to present the numb of likes for each post
			likeCommentdata.Scan(&AllData.Postpage.Comments[i].Likes)
			dislikeCommentdata := DB.QueryRow("SELECT COUNT(user_id) FROM interaction_comments where comment_id = ? AND interaction = ?", AllData.Postpage.Comments[i].Comment_id, 0) // to present the numb of dislikes for each post
			dislikeCommentdata.Scan(&AllData.Postpage.Comments[i].Dislikes)
		}
	}
}
