package forum

import (
	model "forum/model"
	"html/template"
	"net/http"
)

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
	model.LiveSession = model.EmptySession
	err = model.GetUserPostInteractions()
	if err != nil {
		// XXX Currently set to return nil only to allow logout
	}
	t.ExecuteTemplate(w, "index.html", model.AllData)
}
