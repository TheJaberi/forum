package forum

import (
	forum "forum/functions"
	"html/template"
	"net/http"
)

func MainHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
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
	forum.ViewCategory()
	forum.ViewPosts()
	forum.AllData.AllPosts = RSort(forum.AllPosts)
	forum.AllData.AllCategories = forum.AllCategories
	forum.AllData.CategoryCheck = true
	// forum.AllData.LoggedUser = forum.LoggedUser
	// forum.AllData.LoggedUserID = forum.LoggedUser.Userid
	// forum.AllData.IsLogged = forum.LoggedUser.Registered
	t.ExecuteTemplate(w, "index.html", forum.AllData)
}

func RSort(list []forum.Post) []forum.Post {
	var arrAllPosts []forum.Post
	for i:=len(list)-1;i>0;i--{
		arrAllPosts = append(arrAllPosts, list[i])
	}
	return arrAllPosts
}