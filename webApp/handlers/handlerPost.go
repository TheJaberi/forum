package forum

import (
	"fmt"
	forum "forum/functions"
	"html/template"
	"net/http"
)
func Posthandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("test1")
	if req.URL.Path != "/post" {
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
	title := req.FormValue("title")
	body := req.FormValue("post")
	forum.CreatePost(title, body)
	t.ExecuteTemplate(w, "main.html", nil)
}
 