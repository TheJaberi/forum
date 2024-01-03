package forum

import (
	"context"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InsertInteraction(post_id int, user_id int, likeOrDislike int) {
	query := "INSERT INTO `Interaction` (`post_id`, `user_id`, `interaction`) VALUES (?, ?, ?)"
	_, err2 := DB.ExecContext(context.Background(), query, post_id, user_id, likeOrDislike)
	if err2 != nil { // the post is added using the ExecContext along with the userid which is in the LoggedUser variable
		log.Fatal(err2)
	}

}
