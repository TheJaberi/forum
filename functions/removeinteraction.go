package forum

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func RemoveInteraction(post_id int, user_id int, post bool, comment_id int) {
	if post {
	query := "DELETE FROM `Interaction` where post_id = ? AND user_id = ?"
	_, err2 := DB.ExecContext(context.Background(), query, post_id, user_id)
	if err2 != nil { // the post is added using the ExecContext along with the userid which is in the LoggedUser variable
		log.Fatal(err2)
	}
} else {
	query := "DELETE FROM `interaction_comments` where comment_id = ?, post_id = ? AND user_id = ?"
	_, err2 := DB.ExecContext(context.Background(), query, comment_id, post_id, user_id)
	if err2 != nil { // the post is added using the ExecContext along with the userid which is in the LoggedUser variable
		log.Fatal(err2)
	}
}
}
