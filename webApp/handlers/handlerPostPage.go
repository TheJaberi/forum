package forum

import (
	forum "forum/functions"
	"html/template"
	"net/http"
	"strconv"
)

func HandlerPostPage(w http.ResponseWriter, req *http.Request) {
	var postData forum.Post
	if req.URL.Path != "/postpage/" {
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
	postNumb, _ := strconv.Atoi(req.URL.Query().Get("id"))
	postData = forum.AllPosts[postNumb-1]
	postData.LoggedUser = forum.LoggedUser.Registered
	t.ExecuteTemplate(w, "postpage.html", postData)
}
