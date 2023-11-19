package forum

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func RemoveInteraction(post_id int, user_id int) {
	Database, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer Database.Close()
	query := "DELETE FROM `Interaction` where (`post_id`, `user_id`) VALUES (?, ?)"
	_, err2 := Database.ExecContext(context.Background(), query, post_id, user_id)
	if err2 != nil { // the post is added using the ExecContext along with the userid which is in the LoggedUser variable
		log.Fatal(err2)
	}

}
