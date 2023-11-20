package forum

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func UpdateInteraction(post_id int, user_id int, likeOrDislike int){
	Database, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer Database.Close()
	query := "UPDATE Interaction SET interaction = ? where post_id= ? AND user_id = ?"
	_, err2 := Database.ExecContext(context.Background(), query, likeOrDislike, post_id, user_id)
	if err2 != nil { // the post is added using the ExecContext along with the userid which is in the LoggedUser variable
		log.Fatal(err2)
	}
}