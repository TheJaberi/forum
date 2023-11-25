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
	forum.AllData.AllPosts = forum.AllPosts
	forum.AllData.AllCategories = forum.AllCategories
	forum.AllData.LoggedUser = forum.LoggedUser
	forum.AllData.IsLogged = false
	t.ExecuteTemplate(w, "main.html", forum.AllData)
}
