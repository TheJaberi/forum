package forum
import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)
func UpdatePosts(){
	Database, errdatabase := sql.Open("sqlite3", "./forum.db")
	if errdatabase != nil {
		log.Fatal(errdatabase)
	}
	defer Database.Close()
	if LoggedUser.Registered{ // if the user is logged in the fact that he has liked or disliked the post is saved in all posts
		for i:= 0;i<len(AllPosts);i++{
			var interaction int
			postData := Database.QueryRow("SELECT interaction from Interaction where post_id = ? AND user_id = ?", i+1, LoggedUser.Userid)
			errpost := postData.Scan(&interaction)
			if errpost!=nil{
				continue
			} else {
				if interaction==1{
					AllPosts[i].Userlike = true
				} else {
					AllPosts[i].UserDislike = true
				}
			}
		}
	}
}