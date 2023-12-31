package forum

import (
	forum "forum/functions"
	"html/template"
	"net/http"
)

func HandlerLogin(w http.ResponseWriter, req *http.Request) {
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
	email := req.FormValue("email") // when the login button is clicked the username data is assigned to a variable
	password := req.FormValue("password") // when the login button is clicked the password data is assigned to a variable
	forum.UserDbLogin(email, password)       // login func goes over all the rows in the users table and checks if it matches
	forum.UpdatePosts()
	t.ExecuteTemplate(w, "index.html", forum.AllData)
}
