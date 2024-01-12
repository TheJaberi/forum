package forum

import (
	forum "forum/functions"
	"html/template"
	"net/http"
	"strconv"
)

func HandlerComments(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/comment" {
		ErrorHandler(w, req, http.StatusNotFound)
		return
	}
	if req.Method != "POST" {
		ErrorHandler(w, req, http.StatusMethodNotAllowed)
		return
	}
	t, err := template.ParseFiles(HTMLs...)
	if err != nil {
		ErrorHandler(w, req, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	post_id, _ := strconv.Atoi(req.FormValue("postid")) // when the createpost button is clicked the title data is assigned to a variable
	commentContent := req.FormValue("commentContent")
	commentContent = forum.AdjustText(commentContent)
	forum.CreateComment(commentContent, post_id) // create post adds the title and body to the table in the database
	forum.ViewPosts()
	postData := forum.AllData.AllPosts[post_id-1]
	postData.LoggedUser = true
	t.ExecuteTemplate(w, "postpage.html", postData)
}
