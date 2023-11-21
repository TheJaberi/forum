package forum

import (
	forum "forum/functions"
	"html/template"
	"net/http"
	"strconv"
)

func HandlerMyFilter(w http.ResponseWriter, req *http.Request) {
	var filteredPosts []forum.Post
	if req.URL.Path != "/mylikes/" && req.URL.Path != "/myposts/" && req.URL.Path != "/mydislikes/"{
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
	user_id, _ := strconv.Atoi(req.FormValue("interaction"))
	for i := 0; i < len(forum.AllPosts); i++ {
		if req.URL.Path == "/myposts/" && forum.AllPosts[i].UserID == user_id {
			filteredPosts = append(filteredPosts, forum.AllPosts[i])
		}
		if req.URL.Path == "/mylikes/" && forum.AllPosts[i].Userlike {
			filteredPosts = append(filteredPosts, forum.AllPosts[i])
		}
		if req.URL.Path == "/mydislikes/" && forum.AllPosts[i].UserDislike {
			filteredPosts = append(filteredPosts, forum.AllPosts[i])
		}
	}
	forum.AllData.AllPosts = filteredPosts
	forum.AllData.AllCategories = forum.AllCategories
	t.ExecuteTemplate(w, "main.html", forum.AllData)
}
