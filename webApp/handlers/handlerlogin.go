package forum

import (
	"fmt"
	forum "forum/functions"
	"html/template"
	"net/http"
)

func HandlerLogin(w http.ResponseWriter, req *http.Request){
	fmt.Println("test1")
	if req.URL.Path != "/login" {
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
	username := req.FormValue("username")
	password := req.FormValue("password")
	forum.Login(username, password)
	t.ExecuteTemplate(w, "main.html", nil)
}