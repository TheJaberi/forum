package forum

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"os"
	// "fmt"
	forum "forum/functions"

	_ "github.com/mattn/go-sqlite3"
)

func HandlerCommentsLikes(w http.ResponseWriter, req *http.Request) {
	// if !forum.AllData.IsLogged {
	// 	fmt.Println(11)
	// 	ErrorHandler(w, req, http.StatusNotFound)
	// 	return
	// }
	if req.URL.Path != "/commentlike/" && req.URL.Path != "/commentdislike/" {
		ErrorHandler(w, req, http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		ErrorHandler(w, req, http.StatusMethodNotAllowed)
		return
	}
	t, err := template.ParseFiles(HTMLs...)
	if err != nil {
		ErrorHandler(w, req, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	addComment_id, _ := strconv.Atoi(req.FormValue("commentInteraction")) // comment interaction handles the data from like or dislike button if the user logged hasn't already clicked on it
	fmt.Println(addComment_id)
	remComment_id, _ := strconv.Atoi(req.FormValue("removeInteraction"))  // remove interaction handles the data from like or dislike button if the user logged has already clicked on it
	fmt.Println(remComment_id)
	user_id := forum.LoggedUser.Userid
	if addComment_id > remComment_id { // which ever value is greater determines whether to add or remove
		os.Exit(0)
		if req.URL.Path == "/likecomment/" {
			if !forum.AllData.Postpage.Comments[addComment_id-1].CommentUserDislike {
				forum.InsertInteraction(forum.AllData.Postpage.PostID, user_id, 1, false, addComment_id) // insert adds the interaction to the database 1 is like 0 is dislike
				forum.AllData.Postpage.Comments[addComment_id-1].CommentUserlike = true                         // changes the comment like or dislike for the logged in user in the all comments var
			} else {
				forum.UpdateInteraction(forum.AllData.Postpage.PostID, user_id, 1, false, addComment_id) // update is used if a like has to be changed to a dislike or vice versa
				forum.AllData.Postpage.Comments[addComment_id-1].CommentUserlike = true
				forum.AllData.Postpage.Comments[addComment_id-1].CommentUserDislike = false
			}
		} else {
			if !forum.AllData.Postpage.Comments[addComment_id-1].CommentUserlike {
				forum.InsertInteraction(forum.AllData.Postpage.PostID, user_id, 0, false, addComment_id)
				forum.AllData.Postpage.Comments[addComment_id-1].CommentUserDislike = true
			} else {
				forum.UpdateInteraction(forum.AllData.Postpage.PostID, user_id, 0, false, addComment_id)
				forum.AllData.Postpage.Comments[addComment_id-1].CommentUserDislike = true
				forum.AllData.Postpage.Comments[addComment_id-1].CommentUserlike = false
			}
		}
	} else {
		forum.RemoveInteraction(forum.AllData.Postpage.PostID, user_id, false, remComment_id) // remove is greater means there is already an interaction that needs to be removed
		forum.AllData.Postpage.Comments[remComment_id-1].CommentUserlike = false
		forum.AllData.Postpage.Comments[remComment_id-1].CommentUserDislike = false
	}
	fmt.Println(11)
	forum.AllData.Postpage.LoggedUser = true
	t.ExecuteTemplate(w, "postpage.html", forum.AllData.Postpage)
}
