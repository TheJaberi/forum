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
	username := req.FormValue("username") // when the login button is clicked the username data is assigned to a variable
	password := req.FormValue("password") // when the login button is clicked the password data is assigned to a variable
	forum.Login(username, password)       // login func goes over all the rows in the users table and checks if it matches
	t.ExecuteTemplate(w, "main.html", forum.AllData)
}
