package forum

import (
	"fmt"
	model "forum/model"
	"html/template"
	"net/http"
)

// handles the logout process
func HandlerLogout(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/logout/" {
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
	model.AllData.LoggedUser = model.Empty
	model.AllData.IsLogged = false
	model.AllData.TypeAdmin = false
	fmt.Println(model.LiveSession)
	fmt.Println(model.LoginCookie)
	model.LoginCookie = &http.Cookie{
		Name:     model.LiveSession.Name,
		Value:    model.LiveSession.Uuid.String(),
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	fmt.Println(model.LoginCookie)
	http.SetCookie(w, model.LoginCookie)
	cookie, err := req.Cookie(model.LiveSession.Name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cookie)
	model.LiveSession = model.EmptySession
	fmt.Println(model.LiveSession)
	t.ExecuteTemplate(w, "index.html", model.AllData)
}
