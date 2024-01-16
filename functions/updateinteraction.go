package forum

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func UpdateInteraction(post_id int, user_id int, likeOrDislike int, post bool, comment_id int){
	if post {
	query := "UPDATE Interaction SET interaction = ? where post_id= ? AND user_id = ?"
	_, err2 := DB.ExecContext(context.Background(), query, likeOrDislike, post_id, user_id)
	if err2 != nil { // the post is added using the ExecContext along with the userid which is in the LoggedUser variable
		log.Fatal(err2)
	}
} else {
	query := "UPDATE interaction_comments SET interaction = ? where comment_id= ? AND user_id = ?"
	_, err2 := DB.ExecContext(context.Background(), query, likeOrDislike, comment_id, user_id)
	if err2 != nil { // the post is added using the ExecContext along with the userid which is in the LoggedUser variable
		log.Fatal(err2)
	}
}
}