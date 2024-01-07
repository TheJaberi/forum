package forum

import (
	forumVar "forum"
	forum "forum/functions"
	"html/template"
	"net/http"
)

func HandlerRegister(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/register" {
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
	var NewApplicant forumVar.Applicant
	NewApplicant.Username = req.FormValue("username")         // when the register button is clicked the username data is assigned to a variable
	NewApplicant.Password = []byte(req.FormValue("password")) // when the register button is clicked the password data is assigned to a variable
	NewApplicant.Email = req.FormValue("email")
	NewApplicant.Type = "Member"

	forum.UserDbRegisteration(NewApplicant, forum.DB) // NewUser adds the username and password to the database
	t.ExecuteTemplate(w, "index.html", nil)
}
