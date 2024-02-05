package forum

import (
	"html/template"
	"net/http"
	"strconv"
	forum "forum/functions"
	_ "github.com/mattn/go-sqlite3"
)

func HandlerCommentsLikes(w http.ResponseWriter, req *http.Request) {
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
	remComment_id, _ := strconv.Atoi(req.FormValue("removeInteraction"))  // remove interaction handles the data from like or dislike button if the user logged has already clicked on it
	user_id := forum.LoggedUser.Userid
	commentPos := 0
	for i:=0;i<len(forum.AllData.Postpage.Comments);i++{
		if forum.AllData.Postpage.Comments[i].Comment_id == addComment_id || forum.AllData.Postpage.Comments[i].Comment_id == remComment_id {
			commentPos = i
			break
		}
	}
	if addComment_id > remComment_id { // which ever value is greater determines whether to add or remove
		if req.URL.Path == "/commentlike/" {
			if !forum.AllData.Postpage.Comments[commentPos].CommentUserDislike {
				forum.InsertInteraction(forum.AllData.Postpage.PostID, user_id, 1, false, addComment_id) // insert adds the interaction to the database 1 is like 0 is dislike
				forum.AllData.Postpage.Comments[commentPos].CommentUserlike = true                         // changes the comment like or dislike for the logged in user in the all comments var
			} else {
				forum.UpdateInteraction(forum.AllData.Postpage.PostID, user_id, 1, false, addComment_id) // update is used if a like has to be changed to a dislike or vice versa
				forum.AllData.Postpage.Comments[commentPos].CommentUserlike = true
				forum.AllData.Postpage.Comments[commentPos].CommentUserDislike = false
			}
		} else {
			if !forum.AllData.Postpage.Comments[commentPos].CommentUserlike {
				forum.InsertInteraction(forum.AllData.Postpage.PostID, user_id, 0, false, addComment_id)
				forum.AllData.Postpage.Comments[commentPos].CommentUserDislike = true
			} else {
				forum.UpdateInteraction(forum.AllData.Postpage.PostID, user_id, 0, false, addComment_id)
				forum.AllData.Postpage.Comments[commentPos].CommentUserDislike = true
				forum.AllData.Postpage.Comments[commentPos].CommentUserlike = false
			}
		}
	} else {
		forum.RemoveInteraction(forum.AllData.Postpage.PostID, user_id, false, remComment_id) // remove is greater means there is already an interaction that needs to be removed
		forum.AllData.Postpage.Comments[commentPos].CommentUserlike = false
		forum.AllData.Postpage.Comments[commentPos].CommentUserDislike = false
	}
	forum.AllData.Postpage.LoggedUser = true
	t.ExecuteTemplate(w, "postpage.html", forum.AllData.Postpage)
}
