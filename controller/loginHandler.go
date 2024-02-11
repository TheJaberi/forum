package forum

import (
	"html/template"
	"net/http"

	model "forum/model"
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

	loginCookie, err := model.UserLogin(req.FormValue("email"), req.FormValue("password"))
	if err != nil {
		model.LoginError2 = true
	} else {
		model.LoginError2 = false
	}

	http.SetCookie(w, loginCookie)
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "index.html", model.AllData)
}
