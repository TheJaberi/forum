package forum

import (
	"fmt"
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

	email := req.FormValue("email")                    // when the login button is clicked the username data is assigned to a variable
	password := req.FormValue("password")              // when the login button is clicked the password data is assigned to a variable
	session, err := forum.UserDbLogin(email, password) // login func goes over all the rows in the users table and checks if it matches
	if err != nil {
		fmt.Println(err)
		return
	}
	sessionCookie := &http.Cookie{
		Name:     session.Name,
		Value:    session.Uuid.String(),
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
	}
	// Set cookie in response headers
	fmt.Println(sessionCookie)
	http.SetCookie(w, sessionCookie)
	forum.UpdatePosts()
	w.WriteHeader(http.StatusOK)
	// http.Redirect(w, req, "/", http.StatusOK)
	t.ExecuteTemplate(w, "index.html", forum.AllData)
}
